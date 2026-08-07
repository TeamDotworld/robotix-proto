#!/usr/bin/env python3
"""
Generates Go types for VDA5050 3.0.0 directly from the official JSON schemas
published at https://github.com/VDA5050/VDA5050 (json_schemas/*.schema).

Run:  python3 gen_types.py <schema_dir> <out_file>

Design rules
------------
* required field            -> value type
* optional field            -> pointer (scalars/objects) or nil-able slice/map
* enums                     -> named string type + typed constants
* inline objects            -> hoisted to a named struct <Parent><Field>
* allOf/if-then properties  -> merged in as optional (conditionally required)
"""

import json
import os
import re
import sys
from collections import OrderedDict

# --- naming ----------------------------------------------------------------

INITIALISMS = {
    "id": "ID", "ids": "IDs", "url": "URL", "uri": "URI", "json": "JSON",
    "http": "HTTP", "https": "HTTPS", "api": "API", "ip": "IP", "soc": "SOC",
    "xy": "XY", "vda": "VDA", "mqtt": "MQTT", "agv": "AGV", "nurbs": "NURBS",
}

# Fields that the spec types as unsigned 32-bit integers.
UINT32_FIELDS = {
    "headerId", "sequenceId", "orderUpdateId", "lastNodeSequenceId",
    "degree", "referenceStateHeaderId",
}

# Readable names for structs whose auto-derived name is unwieldy or collides.
# Key = auto-derived name, value = the name actually emitted.
NAME_OVERRIDES = {
    # `zoneSet` means two different things in state.schema and zoneSet.schema.
    "StateZoneSet": "ZoneSetInfo",
    # instantActions carries its own inline action (blockingType is always NONE).
    "InstantActionsAction": "InstantAction",
    "InstantActionActionParameter": "ActionParameter",
    "ResponsesResponse": "Response",
    "NodeNodePosition": "NodePosition",
    "NodePositionAllowedDeviationXY": "AllowedDeviationXY",
    "ActionActionParameter": "ActionParameter",
    "TrajectoryControlPoint": "ControlPoint",
    "IntermediatePathPolyline": "PolylinePoint",
    "LoadBoundingBoxReference": "BoundingBoxReference",
    "LoadLoadDimensions": "LoadDimensions",
    "ErrorErrorReference": "ErrorReference",
    "InfoInfoReference": "InfoReference",
    "FactsheetTypeSpecification": "TypeSpecification",
    "FactsheetPhysicalParameters": "PhysicalParameters",
    "FactsheetProtocolLimits": "ProtocolLimits",
    "FactsheetProtocolLimitsMaximumStringLengths": "MaximumStringLengths",
    "FactsheetProtocolLimitsMaximumArrayLengths": "MaximumArrayLengths",
    "FactsheetProtocolLimitsTiming": "Timing",
    "FactsheetProtocolFeatures": "ProtocolFeatures",
    "FactsheetProtocolFeaturesOptionalParameter": "OptionalParameter",
    "FactsheetProtocolFeaturesMobileRobotAction": "MobileRobotAction",
    "MobileRobotActionActionParameter": "ActionParameterDefinition",
    "FactsheetMobileRobotGeometry": "MobileRobotGeometry",
    "MobileRobotGeometryWheelDefinition": "WheelDefinition",
    "WheelDefinitionPosition": "WheelPosition",
    "MobileRobotGeometryEnvelopes2dItem": "Envelope2D",
    "MobileRobotGeometryEnvelopes2dItemVertice": "Envelope2DVertex",
    "MobileRobotGeometryEnvelopes3dItem": "Envelope3D",
    "ProtocolLimitsMaximumStringLengths": "MaximumStringLengths",
    "ProtocolLimitsMaximumArrayLengths": "MaximumArrayLengths",
    "ProtocolLimitsTiming": "ProtocolTiming",
    "ProtocolFeaturesOptionalParameter": "OptionalParameter",
    "ProtocolFeaturesMobileRobotAction": "MobileRobotAction",
    "ProtocolFeaturesMobileRobotActionActionParameter": "ActionParameterDefinition",
    "IntermediatePathPolylineItem": "PolylinePoint",
    "FactsheetLoadSpecification": "LoadSpecification",
    "LoadSpecificationLoadSet": "LoadSet",
    "LoadSetBoundingBoxReference": "BoundingBoxReference",
    "LoadSetLoadDimensions": "LoadDimensions",
    "FactsheetMobileRobotConfiguration": "MobileRobotConfiguration",
    "MobileRobotConfigurationVersion": "VersionInfo",
    "MobileRobotConfigurationNetwork": "NetworkInfo",
    "MobileRobotConfigurationBatteryCharging": "BatteryChargingInfo",
    # nodeState.nodePosition is a reduced form of order's nodePosition (it carries
    # no allowedDeviation*), so it deliberately gets its own type.
    "NodeStateNodePosition": "NodeStatePosition",
    "StateVelocity": "Velocity",
}

# `zoneSet` is defined by two schemas with incompatible shapes; disambiguate by
# origin schema.  Key = (schema file, definition name).
DEF_RENAMES = {
    ("state.schema", "zoneSet"): "ZoneSetInfo",
    ("zoneSet.schema", "zoneSet"): "ZoneSet",
}


def go_name(name: str) -> str:
    parts = re.findall(r"[A-Z]+(?![a-z])|[A-Z][a-z0-9]*|[a-z0-9]+", name)
    out = []
    for p in parts:
        low = p.lower()
        out.append(INITIALISMS.get(low, p[:1].upper() + p[1:]))
    return "".join(out) or name


def enum_const_name(type_name: str, value: str) -> str:
    return type_name + go_name(value.lower())


# --- generator -------------------------------------------------------------

class Generator:
    def __init__(self):
        self.structs = OrderedDict()   # name -> [(goField, goType, jsonTag, doc, unit, req)]
        self.enums = OrderedDict()     # name -> (values, doc)
        self.order = []
        self.schema = None
        self.schema_file = None
        self.building = set()
        self.emitted_docs = {}

    def rename(self, candidate: str) -> str:
        return NAME_OVERRIDES.get(candidate, candidate)

    def def_name(self, defname: str) -> str:
        override = DEF_RENAMES.get((self.schema_file, defname))
        return override if override else self.rename(go_name(defname))

    # -- schema helpers --
    def resolve(self, node):
        seen = 0
        while isinstance(node, dict) and "$ref" in node:
            ref = node["$ref"]
            assert ref.startswith("#/"), ref
            cur = self.schema
            for part in ref[2:].split("/"):
                cur = cur[part]
            node = cur
            seen += 1
            if seen > 16:
                raise RuntimeError("ref cycle")
        return node

    def ref_name(self, node):
        """Return the Go name for a plain $ref node, else None."""
        if isinstance(node, dict) and "$ref" in node:
            return self.def_name(node["$ref"].split("/")[-1])
        return None

    def merged_props(self, node):
        """Flatten properties + allOf/if-then properties. Returns (props, required)."""
        props = OrderedDict(node.get("properties", {}))
        required = list(node.get("required", []))
        for sub in node.get("allOf", []):
            branch = sub.get("then", sub)
            for k, v in branch.get("properties", {}).items():
                if k not in props:
                    props[k] = v
            # NOTE: conditional 'required' is deliberately NOT promoted --
            # those fields stay optional in Go and are enforced by the
            # JSON-Schema validator at runtime.
        return props, required

    # -- type mapping --
    def go_type(self, node, parent: str, field: str, required: bool):
        ref = self.ref_name(node)
        if ref:
            self.ensure_def(node["$ref"].split("/")[-1])
            return ("*" if not required else "") + ref

        node = self.resolve(node)
        t = node.get("type")

        # value-of-anything (actionParameter.value)
        if isinstance(t, list) or t is None:
            return "any"

        if t == "string":
            if "enum" in node:
                name = self.add_enum(parent, field, node)
                return ("*" if not required else "") + name
            return ("*" if not required else "") + "string"

        if t == "boolean":
            return ("*" if not required else "") + "bool"

        if t == "integer":
            base = "uint32" if field in UINT32_FIELDS or node.get("minimum", -1) >= 0 else "int64"
            return ("*" if not required else "") + base

        if t == "number":
            return ("*" if not required else "") + "float64"

        if t == "array":
            items = node.get("items", {})
            # Prefer the schema's own `title` for the element name so that e.g.
            # `actionParameters[]` becomes ActionParameter rather than
            # ActionActionParametersActionParameter.
            resolved = items if "$ref" in items else self.resolve(items)
            elem_field = resolved.get("title") or singular(field)
            inner = self.go_type(items, parent, elem_field, True)
            return "[]" + inner

        if t == "object":
            if not node.get("properties") and not node.get("allOf"):
                return "map[string]any"
            name = self.add_struct(parent + go_name(field), node)
            return ("*" if not required else "") + name

        raise RuntimeError(f"unhandled type {t} at {parent}.{field}")

    def add_enum(self, parent: str, field: str, node) -> str:
        name = go_name(field)
        if name in ("Type", "Status", "State", "Level"):
            name = go_name(parent) + name
        vals = node["enum"]
        if name in self.enums and self.enums[name][0] != vals:
            name = go_name(parent) + go_name(field)
        self.enums[name] = (vals, node.get("description", ""))
        return name

    def build_fields(self, name: str, node):
        fields = []
        props, required = self.merged_props(node)
        for fname, fnode in props.items():
            req = fname in required
            gtype = self.go_type(fnode, name, fname, req)
            tag = fname if req else fname + ",omitempty"
            resolved = self.resolve(fnode)
            desc = fnode.get("description") or resolved.get("description", "")
            unit = resolved.get("unit")
            fields.append((go_name(fname), gtype, tag, desc, unit, req))
        return fields

    def add_struct(self, name: str, node, is_def: bool = False) -> str:
        name = name if is_def else self.rename(go_name(name))
        if name in self.building:        # recursion guard
            return name
        if name in self.structs:
            self.building.add(name)
            fields = self.build_fields(name, node)
            self.building.discard(name)
            if shape(fields) != shape(self.structs[name]):
                raise RuntimeError(
                    f"struct name collision with differing shape: {name}\n"
                    f"  existing: {shape(self.structs[name])}\n"
                    f"  new     : {shape(fields)}"
                )
            # Same shape, different prose: keep whichever documentation is richer.
            self.structs[name] = [
                new if len(new[3] or "") > len(old[3] or "") else old
                for old, new in zip(self.structs[name], fields)
            ]
            return name
        self.building.add(name)
        fields = self.build_fields(name, node)
        self.building.discard(name)
        self.structs[name] = fields
        self.order.append(name)
        return name

    def ensure_def(self, defname: str):
        node = self.schema.get("definitions", {}).get(defname)
        if node is None:
            node = self.schema.get("$defs", {}).get(defname)
        if node is None:
            raise RuntimeError("missing definition " + defname)
        self.add_struct(self.def_name(defname), node, is_def=True)

    # -- entrypoint --
    def add_schema(self, path: str, root_name: str):
        self.schema = json.load(open(path))
        self.schema_file = os.path.basename(path)
        self.add_struct(root_name, self.schema, is_def=True)
        self.emitted_docs[root_name] = self.schema.get("description", "")


def shape(fields):
    """Structural fingerprint of a struct: name, Go type, JSON tag, requiredness.

    Documentation text and units are excluded so that the same object described
    with different prose in two schemas is still recognised as one type.
    """
    return [(f, t, tag, req) for (f, t, tag, _desc, _unit, req) in fields]


def singular(name: str) -> str:
    if name.endswith("ies"):
        return name[:-3] + "y"
    if name.endswith("sses") or name.endswith("shes"):
        return name[:-2]
    if name.endswith("s") and not name.endswith("ss"):
        return name[:-1]
    return name + "Item"


def wrap_doc(text: str, indent: str = "", width: int = 96):
    if not text:
        return []
    text = " ".join(text.split())
    words, lines, cur = text.split(" "), [], ""
    for w in words:
        if cur and len(cur) + len(w) + 1 > width:
            lines.append(cur)
            cur = w
        else:
            cur = (cur + " " + w).strip()
    if cur:
        lines.append(cur)
    return [f"{indent}// {ln}" for ln in lines]


def main():
    schema_dir, out_file = sys.argv[1], sys.argv[2]

    # root name per schema file
    roots = [
        ("order.schema", "Order"),
        ("instantActions.schema", "InstantActions"),
        ("state.schema", "State"),
        ("visualization.schema", "Visualization"),
        ("connection.schema", "Connection"),
        ("factsheet.schema", "Factsheet"),
        ("zoneSet.schema", "ZoneSetMessage"),
        ("responses.schema", "Responses"),
    ]

    g = Generator()
    for fname, root in roots:
        g.add_schema(os.path.join(schema_dir, fname), root)

    out = []
    w = out.append
    w("// Code generated by gen_types.py from the official VDA5050 3.0.0 JSON")
    w("// schemas (https://github.com/VDA5050/VDA5050). DO NOT EDIT.")
    w("//")
    w("// Conventions:")
    w("//   - Required fields are value types.")
    w("//   - Optional fields are pointers (or nil-able slices) with `omitempty`,")
    w("//     so that an absent field is never serialised as a zero value.")
    w("//   - Enumerations are named string types with typed constants.")
    w("")
    w("package vda5050")
    w("")

    # enums
    w("// ---------------------------------------------------------------------------")
    w("// Enumerations")
    w("// ---------------------------------------------------------------------------")
    w("")
    for name, (vals, desc) in g.enums.items():
        for ln in wrap_doc(f"{name} " + (desc or "enumeration defined by VDA5050 3.0.0.")):
            w(ln)
        w(f"type {name} string")
        w("")
        w("const (")
        for v in vals:
            w(f'\t{enum_const_name(name, v)} {name} = "{v}"')
        w(")")
        w("")
        # IsValid helper
        w(f"// IsValid reports whether v is one of the values defined by the specification.")
        w(f"func (v {name}) IsValid() bool {{")
        w("\tswitch v {")
        w("\tcase " + ", ".join(enum_const_name(name, v) for v in vals) + ":")
        w("\t\treturn true")
        w("\t}")
        w("\treturn false")
        w("}")
        w("")

    # structs
    w("// ---------------------------------------------------------------------------")
    w("// Messages and objects")
    w("// ---------------------------------------------------------------------------")
    w("")
    for name in g.order:
        fields = g.structs[name]
        doc = g.emitted_docs.get(name, "")
        for ln in wrap_doc(f"{name} " + (doc if doc else f"is the VDA5050 3.0.0 `{name[:1].lower() + name[1:]}` object.")):
            w(ln)
        w(f"type {name} struct {{")
        for fld, gtype, tag, desc, unit, req in fields:
            if desc or unit:
                d = desc or ""
                if unit:
                    d = (d + f" Unit: {unit}.").strip()
                for ln in wrap_doc(d, "\t"):
                    w(ln)
            if not req:
                w("\t// Optional.")
            w(f'\t{fld} {gtype} `json:"{tag}"`')
        w("}")
        w("")

    with open(out_file, "w") as f:
        f.write("\n".join(out))
    print(f"wrote {out_file}: {len(g.structs)} structs, {len(g.enums)} enums, {len(out)} lines")


if __name__ == "__main__":
    main()
