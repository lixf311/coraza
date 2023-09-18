// Copyright 2022 Juan Pablo Tosso and the OWASP Coraza contributors
// SPDX-License-Identifier: Apache-2.0

package engine

import (
	"github.com/lixf311/coraza/v3/testing/profile"
)

var _ = profile.RegisterProfile(profile.Profile{
	Meta: profile.Meta{
		Author:      "jptosso",
		Description: "Test if the chain action works",
		Enabled:     true,
		Name:        "chains.yaml",
	},
	Tests: []profile.Test{
		{
			Title: "chains",
			Stages: []profile.Stage{
				{
					Stage: profile.SubStage{
						Input: profile.StageInput{
							URI: "/test1.php?id=12345",
						},
						Output: profile.ExpectedOutput{
							TriggeredRules:    []int{1, 1313},
							NonTriggeredRules: []int{2, 200, 20},
						},
					},
				},
				{
					Stage: profile.SubStage{
						Input: profile.StageInput{
							URI: "/test2.php?Var2=prepayloadpost",
						},
						Output: profile.ExpectedOutput{
							TriggeredRules:    []int{1, 200},
							NonTriggeredRules: []int{2, 1313, 20},
							LogContains:       "found within ARGS_GET:Var2: prepayloadpost",
						},
					},
				},
				{
					Stage: profile.SubStage{
						Input: profile.StageInput{
							URI: "/testcase3.php?Var3=pre3payloadpost",
						},
						Output: profile.ExpectedOutput{
							TriggeredRules:    []int{1, 300},
							NonTriggeredRules: []int{2, 1313, 20},
							LogContains:       "found within ARGS_GET:Var3: pre3payloadpost",
						},
					},
				},
				{
					Stage: profile.SubStage{
						Input: profile.StageInput{
							URI: "/testcase4.php?Var4=pre4payloadpost",
						},
						Output: profile.ExpectedOutput{
							TriggeredRules:    []int{1, 400},
							NonTriggeredRules: []int{2, 1313, 20},
							LogContains:       "found within ARGS_GET:Var4: pre4payloadpost",
						},
					},
				},
				{
					Stage: profile.SubStage{
						Input: profile.StageInput{
							URI: "/test3.php",
							Headers: map[string]string{
								"Host": "attack20ing.com",
							},
						},
						Output: profile.ExpectedOutput{
							TriggeredRules:    []int{1, 20},
							NonTriggeredRules: []int{2, 1313, 21},
							LogContains:       "FoundChain20 attack20ing.com in REQUEST_HEADERS:Host",
						},
					},
				},
				{
					Stage: profile.SubStage{
						Input: profile.StageInput{
							URI: "/test4.php",
							Headers: map[string]string{
								"Host": "attack21ing.com",
							},
						},
						Output: profile.ExpectedOutput{
							TriggeredRules:    []int{1, 21},
							NonTriggeredRules: []int{20, 2, 1313},
							LogContains:       "FoundSubChain21 REQUEST_HEADERS:Host in MATCHED_VARS_NAMES",
						},
					},
				},
			},
		},
	},
	Rules: `
SecAction "id: 1, log, chain"
  SecAction "msg:'chain 2',chain"
  SecAction "msg:'chain 3',chain"
  SecAction "msg:'chain 4'"
      
SecAction "id: 2, log, chain"
  SecAction "chain"
  SecAction "chain"
  SecRule ARGS "@noMatch" ""

SecRule REQUEST_URI "@rx (\d\d+)" "id:1313, chain, log"
  SecRule REQUEST_METHOD "GET" ""

SecRule ARGS_GET "@rx prepayloadpost"  "id:200, phase:2, log, msg:'Rule Parent 200', \
  logdata:'Matched Data: %{TX.0} found within %{TX.200_MATCHED_VAR_NAME}: %{MATCHED_VAR}',\
  setvar:'tx.200_matched_var_name=%{MATCHED_VAR_NAME}',\
  chain"
  SecRule MATCHED_VAR "@rx pre" "chain"
    SecRule MATCHED_VAR "@rx post"

SecRule ARGS_GET:var3 "@rx pre3payloadpost"  "id:300, phase:2, log, msg:'Rule Parent 300', \
  logdata:'Matched Data: %{TX.0} found within %{TX.300_MATCHED_VAR_NAME}: %{MATCHED_VAR}',\
  setvar:'tx.300_matched_var_name=%{MATCHED_VAR_NAME}',\
  chain"
  SecRule MATCHED_VAR "@rx pre" "chain"
    SecRule MATCHED_VAR "@rx post"

SecRule ARGS_GET:Var4 "@rx pre4payloadpost"  "id:400, phase:2, log, msg:'Rule Parent 400', \
  logdata:'Matched Data: %{TX.0} found within %{TX.400_MATCHED_VAR_NAME}: %{MATCHED_VAR}',\
  setvar:'tx.400_matched_var_name=%{MATCHED_VAR_NAME}',\
  chain"
  SecRule MATCHED_VAR "@rx pre" "chain"
    SecRule MATCHED_VAR "@rx post"

SecRule REQUEST_HEADERS "@rx attack20" \
  "id:20,\
  phase:1,\
  log,\
  msg:'Chained rule Parent test',\
  logdata:'FoundChain20 %{MATCHED_VAR} in %{MATCHED_VAR_NAME}',\
  chain"
  SecRule MATCHED_VARS_NAMES "@rx Host" \
    "chain"
    SecRule REQUEST_HEADERS:Host "@rx attack20" \
      "block"

SecRule REQUEST_HEADERS "@rx attack21" \
  "id:21,\
  phase:1,\
  log,\
  chain"
  SecRule MATCHED_VARS_NAMES "@rx (?i:host)" \
    "msg:'Sub Chain Rule',\
    logdata:'FoundSubChain21 %{MATCHED_VAR} in %{MATCHED_VAR_NAME}',\
    chain"
    SecRule MATCHED_VAR_NAME "@rx MATCHED_VARS_NAMES" \
    "msg:'Sub Sub Chain Rule',\
    logdata:'FoundSubSubChain21 %{MATCHED_VAR} in %{MATCHED_VAR_NAME}',\
    block"
`,
})

var _ = profile.RegisterProfile(profile.Profile{
	Meta: profile.Meta{
		Author:      "M4tteoP",
		Description: "Tests the number of expected matches against a chained rule",
		Enabled:     true,
		Name:        "chains_counter.yaml",
	},
	Tests: []profile.Test{
		{
			Title: "chains_number_matches",
			Stages: []profile.Stage{
				{
					Stage: profile.SubStage{
						Input: profile.StageInput{
							URI:    "/test",
							Method: "POST",
							Headers: map[string]string{
								"Content-type":   "application/x-www-form-urlencoded",
								"custom_header1": "test",
							},
							Data: "test",
						},
						Output: profile.ExpectedOutput{
							TriggeredRules: []int{1, 2},
						},
					},
				},
			},
		},
	},
	Rules: `
SecDebugLogLevel 9
SecRequestBodyAccess On

SecRule REQUEST_URI "test" "id:1, phase:2, t:none, pass, log, chain"
	SecRule REQUEST_URI|REQUEST_BODY|REQUEST_HEADERS:custom_header1 "test" "setvar:'tx.counter=+1'"
	SecRule TX:counter "@eq 3" "id:2, phase:2, t:none, pass, log"
`,
})

var _ = profile.RegisterProfile(profile.Profile{
	Meta: profile.Meta{
		Author:      "M4tteoP",
		Description: "Tests the number of expected matches against a chained rule",
		Enabled:     true,
		Name:        "chains_counter_2.yaml",
	},
	Tests: []profile.Test{
		{
			Title: "chains_number_matches",
			Stages: []profile.Stage{
				{
					Stage: profile.SubStage{
						Input: profile.StageInput{
							URI:    "/test",
							Method: "POST",
							Headers: map[string]string{
								"Content-type":   "application/x-www-form-urlencoded",
								"custom_header1": "test",
							},
							Data: "test",
						},
						Output: profile.ExpectedOutput{
							TriggeredRules: []int{1, 2},
						},
					},
				},
			},
		},
	},
	Rules: `
SecDebugLogLevel 9
SecRequestBodyAccess On

SecRule REQUEST_URI|REQUEST_HEADERS "test" "id:1, phase:2, t:none, pass, log, chain"
    SecRule REQUEST_URI|REQUEST_HEADERS "test" "chain"
        SecRule REQUEST_BODY|REQUEST_HEADERS:custom_header1|REQUEST_HEADERS "test" "setvar:'tx.counter=+1'"
SecRule TX:counter "@eq 3" "id:2, phase:2, t:none, pass, log"
`,
})
