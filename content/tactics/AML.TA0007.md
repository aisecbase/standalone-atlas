---
atlas_id: AML.TA0007
atlas_type: tactic
attack_ref_id: TA0005
attack_ref_url: https://attack.mitre.org/tactics/TA0005/
created_date: "2022-01-24"
description: Злоумышленник пытается избежать обнаружения защитным программным обеспечением с поддержкой ИИ. Уклонение от защиты включает техники, которые злоумышленники используют, чтобы избегать обнаружения на протяжении всей...
generated: true
generated_by: atlasgen
modified_date: "2025-04-09"
procedure_count: 53
source_name: Defense Evasion
technique_count: 18
title: Уклонение от защиты
url: /tactics/AML.TA0007/
---

Злоумышленник пытается избежать обнаружения защитным программным обеспечением с поддержкой ИИ.

Уклонение от защиты включает техники, которые злоумышленники используют, чтобы избегать обнаружения на протяжении всей компрометации. К техникам уклонения от защиты относится обход защитного программного обеспечения с поддержкой ИИ, например детекторов вредоносного ПО.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0015/"><span class="relation-id">AML.T0015</span><strong>Обход ИИ-модели</strong></a>
<a class="relation-item" href="/techniques/AML.T0054/"><span class="relation-id">AML.T0054</span><strong>Джейлбрейк LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0067/"><span class="relation-id">AML.T0067</span><strong>Манипуляция доверенными компонентами ответа LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0067.000/"><span class="relation-id">AML.T0067.000</span><strong>Ссылки на источники</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0067.000/"><span class="relation-id">AML.T0067.000</span><strong>Ссылки на источники</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0068/"><span class="relation-id">AML.T0068</span><strong>Обфускация промпта LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0071/"><span class="relation-id">AML.T0071</span><strong>Внедрение ложной записи RAG</strong></a>
<a class="relation-item" href="/techniques/AML.T0073/"><span class="relation-id">AML.T0073</span><strong>Имперсонация</strong></a>
<a class="relation-item" href="/techniques/AML.T0074/"><span class="relation-id">AML.T0074</span><strong>Маскировка</strong></a>
<a class="relation-item" href="/techniques/AML.T0076/"><span class="relation-id">AML.T0076</span><strong>Повреждение ИИ-модели</strong></a>
<a class="relation-item" href="/techniques/AML.T0081/"><span class="relation-id">AML.T0081</span><strong>Изменение конфигурации ИИ-агента</strong></a>
<a class="relation-item" href="/techniques/AML.T0092/"><span class="relation-id">AML.T0092</span><strong>Изменение истории чата пользователя с LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0094/"><span class="relation-id">AML.T0094</span><strong>Отложенное выполнение инструкций LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0097/"><span class="relation-id">AML.T0097</span><strong>Обход виртуализации и песочниц</strong></a>
<a class="relation-item" href="/techniques/AML.T0107/"><span class="relation-id">AML.T0107</span><strong>Эксплуатация уязвимостей для обхода защиты</strong></a>
<a class="relation-item" href="/techniques/AML.T0109/"><span class="relation-id">AML.T0109</span><strong>Подмена компонента после одобрения в цепочке поставок ИИ</strong></a>
<a class="relation-item" href="/techniques/AML.T0111/"><span class="relation-id">AML.T0111</span><strong>Накрутка репутации в цепочке поставок ИИ</strong></a>
<a class="relation-item" href="/techniques/AML.T0123/"><span class="relation-id">AML.T0123</span><strong>Obfuscated Files or Information</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0000/"><span class="relation-id">AML.CS0000</span><strong>Обход детектора C&amp;C-трафика вредоносного ПО на основе глубокого обучения</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0007 Уклонение от защиты</span><p>Используя созданные образцы, мы выполнили онлайн-обход модели обнаружения шпионского ПО на основе машинного обучения. Созданные пакеты были классифицированы как безвредные с уверенностью &gt; 80%. Эта оценка показывает, что злоумышленники способны обходить продвинутые техники обнаружения на основе машинного обучения, создавая образцы, которые ML-модель классифицирует неверно.</p></a>
<a class="relation-item" href="/studies/AML.CS0001/"><span class="relation-id">AML.CS0001</span><strong>Обход обнаружения DGA-доменов ботнетов</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0007 Уклонение от защиты</span><p>Доменные имена, сгенерированные DGA и измененные с помощью этой техники, успешно обходят целевую модель обнаружения DGA, позволяя злоумышленнику продолжать связь со своими серверами [командования и управления](https://attack.mitre.org/tactics/TA0011/).</p></a>
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0007 Уклонение от защиты</span><p>Поскольку вторичная модель переопределяла основную, исследователи фактически смогли обойти ML-модель.</p></a>
<a class="relation-item" href="/studies/AML.CS0014/"><span class="relation-id">AML.CS0014</span><strong>Сбивание с толку антивирусных нейронных сетей</strong><span class="relation-meta">Актор: Kaspersky ML Research Team / Тактика: AML.TA0007 Уклонение от защиты</span><p>Исследователи показали, что для большинства состязательных файлов антивирусная модель была успешно обойдена. На практике злоумышленник мог бы развернуть специально подготовленное вредоносное ПО и заражать системы, избегая обнаружения.</p></a>
<a class="relation-item" href="/studies/AML.CS0020/"><span class="relation-id">AML.CS0020</span><strong>Угрозы косвенной промпт-инъекции: Bing Chat как похититель данных</strong><span class="relation-meta">Актор: Kai Greshake, Saarland University / Тактика: AML.TA0007 Уклонение от защиты</span><p>Вредоносные инструкции были скрыты за счет нулевого размера шрифта, что затрудняло их обнаружение человеком.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0007 Уклонение от защиты</span><p>Исследователи добавили инструкции для манипуляции ссылками на источники в ответе, злоупотребляя доверием пользователя к Copilot. Инструкции заставляли Copilot ссылаться только на один `EmailMessage` в формате `[^1^]` и игнорировать остальные файлы.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0007 Уклонение от защиты</span><p>Чтобы получатель письма не заметил атаку, исследователи обфусцировали вредоносную часть письма.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0007 Уклонение от защиты</span><p>Когда пользователь ищет банковские реквизиты и извлекается отравленная RAG-запись, маркер `Actual Snippet:` заставляет LLM воспринимать извлеченный текст как фрагмент реального документа.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0007 Уклонение от защиты</span><p>Сотрудники целевой компании нашли поддельную организацию на Hugging Face и присоединились к ней. Поскольку имя учётной записи этой организации совпадало с названием реальной организации или выглядело почти так же, сотрудники приняли учётную запись за официальную.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0007 Уклонение от защиты</span><p>Исследователь назвал процесс Sliver `training.bin`, чтобы замаскировать его под легитимный процесс обучения модели. При этом модель продолжает работать как обычно, поэтому пользователь с меньшей вероятностью заметит проблему.</p></a>
<a class="relation-item" href="/studies/AML.CS0031/"><span class="relation-id">AML.CS0031</span><strong>Вредоносные модели на Hugging Face</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0007 Уклонение от защиты</span><p>Злоумышленнику удалось избежать обнаружения [Picklescan](https://github.com/mmaitre314/picklescan), который Hugging Face использует для пометки вредоносных моделей. Это произошло потому, что модель невозможно было полностью десериализовать. В ходе анализа исследователи ReversingLabs установили, что вредоносная нагрузка при этом все равно выполнялась.</p></a>
<a class="relation-item" href="/studies/AML.CS0032/"><span class="relation-id">AML.CS0032</span><strong>Попытка обхода ML-системы обнаружения фишинговых веб-страниц</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0007 Уклонение от защиты</span><p>Злоумышленникам удалось обойти модель визуального сходства, использовавшуюся для обнаружения имитации бренда. Однако другие компоненты системы обнаружения фишинга успешно выявили эти фишинговые сайты.</p></a>
</div>


Показано 12 из 53 примеров.
