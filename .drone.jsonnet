local paths = [".", "mnsender", "mnclient", "tabmgr"];
local govers = ["1.10.8", "1.11.5", "1.12"];
local fxvers = ["66.0b9", "66.0b12"];

local TestStep(go, fx, dir) = {
  kind: "pipeline",
  name: dir + "-go" + go + "-fx" + fx,
  workspace: {
    base: "/go",
    path: "src/github.com/raohwork/marionette-go"
  },
  steps: [{
    name: "test",
    image: "ronmi/go-firefox",
    environment: {
      GO_VER: go,
      FX_VER: fx,
    },
    commands: [
      "cd " + dir,
      "run-test.sh go test -v -p 2 -bench . -benchmem -cover",
    ],
  }],
};

local byDir(go, fx) = [
  TestStep(go, fx, x) for x in paths
];

local byFx(go) = std.flattenArrays([
  byDir(go, fx) for fx in fxvers
]);

local byGo() = std.flattenArrays([
  byFx(go) for go in govers
]);

byGo()
