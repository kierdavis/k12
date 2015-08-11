# Root directory for this board (e.g. /eda/boards/alu).
BOARD_DIR = .

# Root directory for all EDA files (/eda).
EDA_DIR = ${BOARD_DIR}/../..

# Subdirectories of board root containing schematics and build files.
SCHEMATICS_DIR = ${BOARD_DIR}/schematics
BUILD_DIR = ${BOARD_DIR}/build
WIRES_DIR = ${BUILD_DIR}/wires

# Subdirectory of EDA root containing gEDA symbol definitions.
SYMS_DIR = ${EDA_DIR}/geda-syms

# Subdirectory of EDA root containing footprint definitions (*.toml) for layout
# processes.
FOOTPRINTS_DIR = ${EDA_DIR}/footprints

# Paths to the EDA tool executables.
TOOLS_DIR = ${EDA_DIR}/tools
BIN_DIR = ${TOOLS_DIR}/bin
PARTLIST_BIN = ${BIN_DIR}/partlist
CHECKSCHEM_BIN = ${BIN_DIR}/checkschem
ROUTE_BIN = ${BIN_DIR}/route
RENDER_BIN = ${BIN_DIR}/render

# Lists of all schematic files and all footprint files.
SCHEMATICS = $(shell find ${SCHEMATICS_DIR} -type f -name '*.sch')
FOOTPRINTS = $(shell find ${FOOTPRINTS_DIR} -type f -name '*.toml')

# Build files.
PARTLIST = ${BUILD_DIR}/parts.txt
NETLIST = ${BUILD_DIR}/netlist.net
INITIAL_LAYOUT = ${BOARD_DIR}/layout.toml
ROUTED_LAYOUT = ${BUILD_DIR}/routed.toml
INITIAL_LAYOUT_SVG = ${BUILD_DIR}/layout.svg
ROUTED_LAYOUT_SVG = ${BUILD_DIR}/routed.svg

# Netlist filename and schematic filenames as absolute paths, for when we run
# gnetlist outside of the current directory.
NETLIST_ABS = $(shell realpath ${NETLIST})
SCHEMATICS_ABS = $(shell realpath ${SCHEMATICS})

# Perform the entire build process, by requiring the parts list and both SVGs to
# be made.
all: ${PARTLIST} ${INITIAL_LAYOUT_SVG} ${ROUTED_LAYOUT_SVG}

# Generate individual wiring diagrams using the "render" tool.
wires: ${RENDER_BIN} ${ROUTED_LAYOUT} ${FOOTPRINTS}
	${RENDER_BIN} -individual-wires -flip -output ${WIRES_DIR} ${ROUTED_LAYOUT} ${FOOTPRINTS}

# Remove the build directory.
clean:
	rm -r ${BUILD_DIR}

# Perform all checks.
check: checkschem

# Create the build directory. Used as a prerequisite for other targets.
${BUILD_DIR}:
	mkdir -p ${BUILD_DIR}

# Build the parts list (build/parts.txt) using the "partlist" tool.
${PARTLIST}: ${BUILD_DIR} ${PARTLIST_BIN} ${SCHEMATICS}
	${PARTLIST_BIN} -sympath ${SYMS_DIR} -output ${PARTLIST} ${SCHEMATICS}

# Check schematic files using the "checkschem" tool.
checkschem: ${CHECKSCHEM_BIN} ${SCHEMATICS}
	${CHECKSCHEM_BIN} ${SCHEMATICS}

# Build the netlist (build/netlist.net) from the schematics using gnetlist.
${NETLIST}: ${BUILD_DIR} ${SCHEMATICS}
	@# Run gnetlist in the eda directory, so that it will find gafrc and hence geda-syms
	cd ${EDA_DIR} && gnetlist -g PCB -o ${NETLIST_ABS} ${SCHEMATICS_ABS}

# Create a routed layout (build/routed.toml) using the "route" tool.
${ROUTED_LAYOUT}: ${BUILD_DIR} ${ROUTE_BIN} ${INITIAL_LAYOUT} ${NETLIST} ${FOOTPRINTS}
	${ROUTE_BIN} -output ${ROUTED_LAYOUT} ${INITIAL_LAYOUT} ${NETLIST} ${FOOTPRINTS}

# Render the unrouted layout as an SVG using the "render" tool.
${INITIAL_LAYOUT_SVG}: ${BUILD_DIR} ${RENDER_BIN} ${INITIAL_LAYOUT} ${FOOTPRINTS}
	${RENDER_BIN} -output ${INITIAL_LAYOUT_SVG} ${INITIAL_LAYOUT} ${FOOTPRINTS}

# Render the routed layout as an SVG using the "render" tool.
${ROUTED_LAYOUT_SVG}: ${BUILD_DIR} ${RENDER_BIN} ${ROUTED_LAYOUT} ${FOOTPRINTS}
	${RENDER_BIN} -output ${ROUTED_LAYOUT_SVG} ${ROUTED_LAYOUT} ${FOOTPRINTS}

.PHONY: all clean check checkschem
