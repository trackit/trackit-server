//   Copyright 2020 MSolution.IO
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package rdsReports

const Template = `
{
	"template": "*-` + IndexSuffix + `",
	"version": 5,
	"mappings": ` + Mappings + `
}
`

const Mappings = `
{
	"rds-report": {
		"properties": {
			"account": {
				"type": "keyword"
			},
			"reportDate": {
				"type": "date"
			},
			"reportType": {
				"type": "keyword"
			},
			"instance": {
				"type": "object",
				"properties": {
					"id": {
						"type": "keyword"
					},
					"availabilityZone": {
						"type": "keyword"
					},
					"type": {
						"type": "keyword"
					},
					"engine": {
						"type": "keyword"
					},
					"allocatedStorage": {
						"type": "integer"
					},
					"multiAZ": {
						"type": "boolean"
					},
					"tags": {
						"type": "nested",
						"properties": {
							"key": {
								"type": "keyword"
							},
							"value": {
								"type": "keyword"
							}
						}
					},
					"costs": {
						"type": "object"
					},
					"stats": {
						"type": "object",
						"properties": {
							"cpu": {
								"type": "object",
								"properties": {
										"average": {
											"type": "double"
										},
										"peak": {
											"type": "double"
										}
								}
							},
							"freeSpace": {
								"type": "object",
								"properties": {
										"minimum": {
											"type": "double"
										},
										"maximum": {
											"type": "double"
										},
										"average": {
											"type": "double"
										}
								}
							}
						}
					}
				}
			}
		},
		"_all": {
			"enabled": false
		},
		"numeric_detection": false,
		"date_detection": false
	}
}`
