package parser

import (
	"fmt"
	"os"
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
						Type: "Heading", HeaderName: "header", Level: 1, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "Text", HeaderName: "", Level: 0, Text: `header`, Url: "", Image: "",
							},
							{
								Type: "Heading", HeaderName: "subheader-1", Level: 2, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Text", HeaderName: "", Level: 0, Text: `subheader 1`, Url: "", Image: "",
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `one to three
1 2 3`, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "Heading", HeaderName: "subheader-2", Level: 2, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Text", HeaderName: "", Level: 0, Text: `subheader 2`, Url: "", Image: "",
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `First item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Second item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Third item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Indented item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Indented item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Fourth item`, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "Heading", HeaderName: "subheader-3", Level: 2, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Text", HeaderName: "", Level: 0, Text: `subheader 3`, Url: "", Image: "",
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `First item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Second item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Third item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Indented item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Indented item`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Fourth item`, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "Heading", HeaderName: "subheader-4", Level: 2, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Text", HeaderName: "", Level: 0, Text: `subheader 4`, Url: "", Image: "",
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Open the file containing the Linux mascot.`, Url: "", Image: "",
											},
										},
									},
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
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `Close the file.`, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "Heading", HeaderName: "subheader-5", Level: 2, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Text", HeaderName: "", Level: 0, Text: `subheader 5`, Url: "", Image: "",
									},
								},
							},
							{
								Type: "Heading", HeaderName: "subheader-6", Level: 2, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Text", HeaderName: "", Level: 0, Text: `subheader 6`, Url: "", Image: "",
									},
								},
							},
						},
					},
					{
						Type: "CodeBlock", HeaderName: "", Level: 0, Text: `Sample text.
`, Url: "", Image: "",
					},
					{
						Type: "CodeBlock", HeaderName: "", Level: 0, Text: `[link](http://example.com)
`, Url: "", Image: "",
					},
					{
						Type: "HorizontalRule", HeaderName: "", Level: 0, Text: `---------------`, Url: "", Image: "",
					},
					{
						Type: "List", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "List", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
											},
											{
												Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
							},
						},
					},
					{
						Type: "List", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "List", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
											},
											{
												Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
							},
						},
					},
					{
						Type: "List", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
							},
							{
								Type: "ListItem", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "",
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

func Test_ParseDoc(t *testing.T) {
	md, err := os.ReadFile("../../data/описание_сущности.md")
	require.NoError(t, err)
	print := func(parent *DocumentItem, item *DocumentItem, depht int) {
		ts := ""
		for i := 0; i < depht; i++ {
			ts = ts + "\t"
		}
		fmt.Printf("%v %#v\r\n", ts, *item)
	}
	d, err := ParseMD(md, print)
	require.NoError(t, err)

	expectedDoc := &Document{
		DocumentItems: []*DocumentItem{
			{
				Type: "Document", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
					{
						Type: "Heading", HeaderName: "описание-сущности", Level: 1, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
							{
								Type: "Text", HeaderName: "", Level: 0, Text: `описание сущности`, Url: "", Image: "",
							},
							{
								Type: "Heading", HeaderName: "служебная-секция", Level: 2, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Text", HeaderName: "", Level: 0, Text: `служебная секция`, Url: "", Image: "",
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `тип документа: описание`, Url: "", Image: "",
											},
										},
									},
								},
							},
							{
								Type: "Heading", HeaderName: "сущность", Level: 2, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
									{
										Type: "Text", HeaderName: "", Level: 0, Text: `сущность`, Url: "", Image: "",
									},
									{
										Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `cущность зовут Фывапрол. сущность - это виртуальное разумное существо.
сущность состоит из исполнителя и памяти.`, Url: "", Image: "",
											},
										},
									},
									{
										Type: "Heading", HeaderName: "исполнитель", Level: 3, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `исполнитель`, Url: "", Image: "",
											},
											{
												Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `исполнитель имеет методы выполнения действий.
исполнитель может выполнять действия.
исполнитель хранит информацию в виде фактов.`, Url: "", Image: "",
													},
												},
											},
										},
									},
									{
										Type: "Heading", HeaderName: "факт", Level: 3, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `факт`, Url: "", Image: "",
											},
											{
												Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `факт состоит из фреймов.`, Url: "", Image: "",
													},
												},
											},
										},
									},
									{
										Type: "Heading", HeaderName: "контекст", Level: 3, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `контекст`, Url: "", Image: "",
											},
											{
												Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `контекст это факт, который группирует другие факты по признакам:`, Url: "", Image: "",
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-1", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 1`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `агент совершающий действие.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-2", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 2`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `местоположение совершения действия.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-3", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 3`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `время совершения действия факта.`, Url: "", Image: "",
															},
														},
													},
												},
											},
										},
									},
									{
										Type: "Heading", HeaderName: "действие", Level: 3, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `действие`, Url: "", Image: "",
											},
											{
												Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `действие - это процесс который изменяет окружение.
выполнение действия - это процесс, состоящий из фазы вызова и фазы выполнения.
вызов выполнения действия - это функциональный вызов, который имеет на входе фрейм специального вида имеющий имя действия и контекст.
контекст - это фрейм специального вида содержащий всю информацию об окружении.
исполнитель, как сущность, находится в окружении которое можно назвать комнатой.`, Url: "", Image: "",
													},
												},
											},
										},
									},
									{
										Type: "Heading", HeaderName: "окружение", Level: 3, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `окружение`, Url: "", Image: "",
											},
											{
												Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Комната имеет пол, стены, потолок.
В комнате стоит стол, и шкаф.
В шкафу лежат книги.
Кроме того, в шкафу находится файлер.
В шкафу стоит телетайп для связи с пользователями.
На столе лежат кубики.
На полу стоит почтовый терминал для отправки или получения документов.`, Url: "", Image: "",
													},
												},
											},
										},
									},
									{
										Type: "Heading", HeaderName: "действия-исполнителя", Level: 3, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `действия исполнителя`, Url: "", Image: "",
											},
											{
												Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Исполнитель может производить следующие действия:`, Url: "", Image: "",
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-1-1", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 1`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Прочитать сообщение из телетайпа.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-2-1", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 2`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Написать сообщение в телетайп.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-3-1", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 3`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Выбрать книгу. (книга по сути то словарь)`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-4", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 4`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Искать слово в книге, вернуть список статей с этим словом.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-5", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 5`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Написать статью со словом в книгу.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-6", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 6`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Посмотреть наличие документов в почтовом терминале.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-7", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 7`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Получить документ из почтового терминала.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-8", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 8`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Отправить документ в почтовый терминал.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-9", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 9`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Посмотреть наличие сообщений в телетайпе.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-10", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 10`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Получить сообщение из телетайппюа.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "пункт-11", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `пункт 11`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Отправить сообщение в телетайп.`, Url: "", Image: "",
															},
														},
													},
												},
											},
										},
									},
									{
										Type: "Heading", HeaderName: "пассивные-объекты", Level: 3, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `пассивные объекты`, Url: "", Image: "",
											},
											{
												Type: "Heading", HeaderName: "пол", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Пол.`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Пол это горизонтальная поверхность. Пол имеет форму квадрата.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "стены", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Стены.`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Стена это вертикальная поверхность. Стена имеет форму квадрата.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "шкаф", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Шкаф.`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Шкаф стоит на полу. Шкаф имеет полки. Полка шкафа это горизонтальная поверхность в габаритах шкафа.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "стол", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Стол.`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Стол это горизонтальная поверхность которая поднята над опорой с помощью ножек. стол имеет столешницу и четыре ножки. столешница это поверхность. ножка это цилиндр. Стол стоит на полу.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "кубики", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Кубики.`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `кубик это куб маленького размера. сто кубиков красного цвета лежат на столе. сто кубиков зеленого цвета лежат на столе.`, Url: "", Image: "",
															},
														},
													},
												},
											},
										},
									},
									{
										Type: "Heading", HeaderName: "активные-объекты", Level: 3, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
											{
												Type: "Text", HeaderName: "", Level: 0, Text: `активные объекты`, Url: "", Image: "",
											},
											{
												Type: "Heading", HeaderName: "книга", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Книга.`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `книга это средства хранения информации. Книга имеет страницы. Страница это поверхность в габаритах книги. Страницы сделана из бумаги. Книга стоит на полке шкафа.
книга имеет метод для сохранения статьи.
книга имеет метод для поиска статьи.
книга имеет метод для чтения статьи.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "телетайп", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Телетайп`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Телетайп позволяет открыть (установить) соединение.
Телетайп позволяет принять сообщение и положить его в очередь.
Телетайп позволяет отправить сообщение.
Телетайп позволяет узнать есть ли сообщение в очереди.
Телетайп позволяет отдать сообщение из очереди если там есть документ.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "почтовый-терминал", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Почтовый терминал`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Почтовый терминал позволяет получить документ и положить его в очередь.
Почтовый терминал позволяет отправить документ.
Почтовый терминал позволяет отдать документ из очереди.
Почтовый терминал позволяет узнать есть ли документы в очереди.
Почтовый терминал позволяет получить документ для отправки.`, Url: "", Image: "",
															},
														},
													},
												},
											},
											{
												Type: "Heading", HeaderName: "часы", Level: 4, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
													{
														Type: "Text", HeaderName: "", Level: 0, Text: `Часы`, Url: "", Image: "",
													},
													{
														Type: "Paragraph", HeaderName: "", Level: 0, Text: ``, Url: "", Image: "", DocumentItems: []*DocumentItem{
															{
																Type: "Text", HeaderName: "", Level: 0, Text: `Часы позволяют получить значение времени.`, Url: "", Image: "",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

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

	var equal func(dc1 *DocumentItem, dc2 *DocumentItem, depht int)
	equal = func(dc1 *DocumentItem, dc2 *DocumentItem, depht int) {
		require.Equal(t, dc1.Type, dc2.Type)
		require.Equal(t, dc1.HeaderName, dc2.HeaderName)
		require.Equal(t, dc1.Level, dc2.Level)
		require.Equal(t, dc1.Text, dc2.Text)
		require.Equal(t, dc1.Url, dc2.Url)
		require.Equal(t, dc1.Image, dc2.Image)
		require.Equal(t, len(dc1.DocumentItems), len(dc2.DocumentItems))
		if len(dc1.DocumentItems) > 0 {
			for i := range dc1.DocumentItems {
				equal(dc1.DocumentItems[i], dc2.DocumentItems[i], depht+1)
			}
		}
	}

	require.Equal(t, len(expectedDoc.DocumentItems), len(d.DocumentItems))
	if len(expectedDoc.DocumentItems) > 0 {
		for i := range d.DocumentItems {
			equal(expectedDoc.DocumentItems[i], d.DocumentItems[i], 0)
		}
	}
}
