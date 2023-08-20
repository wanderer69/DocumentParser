package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getNodeRecur(t *testing.T) {
	var mds = `# header

	Sample text.
	
## subheader 1
	[link](http://example.com)
	
---------------
one to three
1 2 3
	
	
## subheader 2
	
1. First item
2. Second item
3. Third item
	1. Indented item
	2. Indented item
4. Fourth item
	
## subheader 3
	
- First item
- Second item
- Third item
	- Indented item
	- Indented item
- Fourth item
	
## subheader 4
	
1. Open the file containing the Linux mascot.
2. Marvel at its beauty.

	![Tux, the Linux mascot](/assets/images/tux.png)

3. Close the file.

## subheader 5

|------|-------------------|
|      |                   |
|------|-------------------|
|      |                   |


## subheader 6

***

---

_________________
	`

	md := []byte(mds)
	print := func(parent *DocumentItem, item *DocumentItem, depht int) {
		ts := ""
		for i := 0; i < depht; i++ {
			ts = ts + "\t"
		}
		fmt.Printf("%v %#v\r\n", ts, *item)
	}

	expectedDoc := &Document{
		DocumentItems: []*DocumentItem{
			{
				Type: "Document", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
					{
						Type: "Heading", HeaderName: "header", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "Text", HeaderName: "", Level: 0, Text: `header`, Url: "", Image: "",
							},
						},
					},
					{
						Type: "CodeBlock", HeaderName: "", Level: 0, Text: `Sample text.
`, Url: "", Image: "",
					},
					{
						Type: "Heading", HeaderName: "subheader-1", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "Text", HeaderName: "", Level: 0, Text: `subheader 1`, Url: "", Image: "",
							},
						},
					},
					{
						Type: "CodeBlock", HeaderName: "", Level: 0, Text: `[link](http://example.com)
`, Url: "", Image: "",
					},
					{
						Type: "HorizontalRule", HeaderName: "", Level: 0, Text: `---------------`, Url: "", Image: "",
					},
					{
						Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "Text", HeaderName: "", Level: 0, Text: `one to three
1 2 3`, Url: "", Image: "",
							},
						},
					},
					{
						Type: "Heading", HeaderName: "subheader-2", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "Text", HeaderName: "", Level: 0, Text: `subheader 2`, Url: "", Image: "",
							},
						},
					},
					{
						Type: "List", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `First item`, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Second item`, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Third item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "List", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Indented item`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Indented item`, Url: "", Image: "",
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Fourth item`, Url: "", Image: "",
											},
										},
									},
								},
							},
						},
					},
					{
						Type: "Heading", HeaderName: "subheader-3", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "Text", HeaderName: "", Level: 0, Text: `subheader 3`, Url: "", Image: "",
							},
						},
					},
					{
						Type: "List", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `First item`, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Second item`, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Third item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "List", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Indented item`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Indented item`, Url: "", Image: "",
															},
														},
													},
												},
											},
										},
									},
								},
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Fourth item`, Url: "", Image: "",
											},
										},
									},
								},
							},
						},
					},
					{
						Type: "Heading", HeaderName: "subheader-4", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "Text", HeaderName: "", Level: 0, Text: `subheader 4`, Url: "", Image: "",
							},
						},
					},
					{
						Type: "List", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Open the file containing the Linux mascot.`, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Marvel at its beauty.`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
											},
											{
												Type: "Image", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "/assets/images/tux.png", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Tux, the Linux mascot`, Url: "", Image: "",
													},
												},
											},
										},
									},
								},
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Close the file.`, Url: "", Image: "",
											},
										},
									},
								},
							},
						},
					},
					{
						Type: "Heading", HeaderName: "subheader-5", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "Text", HeaderName: "", Level: 0, Text: `subheader 5`, Url: "", Image: "",
							},
						},
					},
					{
						Type: "Table", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "TableBody", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "TableRow", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "TableCell", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
											},
											{
												Type: "TableCell", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
											},
										},
									},
									{
										Type: "TableRow", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "TableCell", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `------`, Url: "", Image: "",
													},
												},
											},
											{
												Type: "TableCell", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `-------------------`, Url: "", Image: "",
													},
												},
											},
										},
									},
									{
										Type: "TableRow", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "TableCell", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
											},
											{
												Type: "TableCell", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
											},
										},
									},
								},
							},
						},
					},
					{
						Type: "Heading", HeaderName: "subheader-6", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "Text", HeaderName: "", Level: 0, Text: `subheader 6`, Url: "", Image: "",
							},
						},
					},
					{
						Type: "HorizontalRule", HeaderName: "", Level: 0, Text: `***`, Url: "", Image: "",
					},
					{
						Type: "HorizontalRule", HeaderName: "", Level: 0, Text: `---`, Url: "", Image: "",
					},
					{
						Type: "HorizontalRule", HeaderName: "", Level: 0, Text: `_________________`, Url: "", Image: "",
					},
				},
			},
		},
	}
	d, err := ParseMD(md, print)
	require.NoError(t, err)

	var prnt func(dc *DocumentItem, depht int)
	prnt = func(dc *DocumentItem, depht int) {
		ts := ""
		for i := 0; i < depht; i++ {
			ts = ts + "\t"
		}
		fmt.Printf("%v{\r\n%v\tType:\"%v\", HeaderName:\"%v\", Level:%v, Text:`%v`, Url:\"%v\", Image:\"%v\"", ts, ts,
			dc.Type, dc.HeaderName, dc.Level, dc.Text, dc.Url, dc.Image)
		if len(dc.DocumentItems) > 0 {
			fmt.Printf(", DocumentItems: []*DocumentItem{\r\n")
			for i := range dc.DocumentItems {
				prnt(dc.DocumentItems[i], depht+2)
			}
			fmt.Printf("\r\n%v\t},\r\n", ts)
		} else {
			fmt.Printf(",\r\n")
		}
		fmt.Printf("%v},\r\n", ts)
	}
	fmt.Printf("{\r\n")
	for i := range d.DocumentItems {
		prnt(d.DocumentItems[i], 1)
	}
	fmt.Printf("},\r\n")
	require.Equal(t, expectedDoc.DocumentItems, d.DocumentItems)
	//require.True(t, false)
}
