local paths = [".", "mnsender", "mnclient", "tabmgr"];
local govers = ["1.10.8", "1.11.5", "1.12"];
local fxvers = ["66.0b9", "66.0b12"];
local tabmgrfx = ["64.0", "65.0"];

local TestTabmgrStep(go, fx) = {
  name: "test-tabmgr-go"+go+"-fx"+fx,
  image: "ronmi/go-firefox",
  environment: {
    GO_VER: go,
    FX_VER: fx,
  },
  commands: [
    "run-test.sh go test -p 2 -run TestTabManager -cover ./tabmgr",
  ],
  volumes: [
    {name: "opt", path: "/opt"},
  ],
};

local TestCmd(dir) = [
  "run-test.sh go test -p 2 -bench . -benchmem -cover ./"+dir,
];

local TestStep(go, fx) = {
  name: "test-go"+go+"-fx"+fx,
  image: "ronmi/go-firefox",
  environment: {
    GO_VER: go,
    FX_VER: fx,
  },
  commands: std.flattenArrays([
    TestCmd(dir) for dir in paths
  ]),
  volumes: [
    {name: "opt", path: "/opt"},
  ],
};

local byFx(go) = [
  TestStep(go, fx) for fx in fxvers
];

local byGo() = std.flattenArrays([
  byFx(go) for go in govers
]);

local TestPipeline() = {
  kind: "pipeline",
  name: "testing",
  workspace: {
    base: "/go",
    path: "src/github.com/raohwork/marionette-go/v3"
  },
  steps: [
  ] + byGo() + [
    TestTabmgrStep(go, fx)
    for go in govers
    for fx in tabmgrfx
  ],
  volumes: [
    {name: "opt", temp: {size_limit: "5g"}},
  ],
};

[
  TestPipeline(),
]
