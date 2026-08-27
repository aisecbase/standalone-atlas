---
atlas_id: AML.T0053
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2023-10-25"
description: Злоумышленники могут использовать свой доступ к ИИ-агенту, чтобы вызывать инструменты, к которым имеет доступ агент. LLM часто подключают к другим сервисам или ресурсам через инструменты, чтобы расширить их...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 12
modified_date: "2026-07-31"
platforms:
    - Agentic AI
procedure_count: 23
source_name: AI Agent Tool Invocation
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0005
    - AML.TA0012
    - AML.TA0015
title: Вызов инструментов ИИ-агента
url: /techniques/AML.T0053/
---

Злоумышленники могут использовать свой доступ к ИИ-агенту, чтобы вызывать инструменты, к которым имеет доступ агент. LLM часто подключают к другим сервисам или ресурсам через инструменты, чтобы расширить их возможности. Инструменты могут включать интеграции с другими приложениями, доступ к публичным или приватным источникам данных и возможность выполнять код.

Это может позволить злоумышленникам выполнять API-вызовы к интегрированным приложениям или сервисам, получая повышенные привилегии в системе. Злоумышленники могут использовать подключенные источники данных для получения чувствительной информации. Они также могут использовать LLM, интегрированную с командным или скриптовым интерпретатором, для выполнения произвольных инструкций.

ИИ-агенты могут быть настроены с доступом к инструментам, которые недоступны пользователям напрямую. Злоумышленники могут злоупотреблять этим, чтобы получить доступ к инструментам, которыми иначе не смогли бы воспользоваться.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
<a class="relation-item" href="/tactics/AML.TA0012/"><span class="relation-id">AML.TA0012</span><strong>Повышение привилегий</strong></a>
<a class="relation-item" href="/tactics/AML.TA0015/"><span class="relation-id">AML.TA0015</span><strong>Латеральное перемещение</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Защитные ограничения (Guardrails) для генеративного ИИ</strong><p>Защитные ограничения могут предотвращать вредоносные входные данные, способные привести к компрометации плагина, и обнаруживать персональные данные в выходных данных модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0021/"><span class="relation-id">AML.M0021</span><strong>Правила и инструкции для генеративного ИИ</strong><p>Инструкции для модели могут предписывать ей отказываться отвечать на небезопасные входные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0022/"><span class="relation-id">AML.M0022</span><strong>Выравнивание модели генеративного ИИ</strong><p>Выравнивание модели может повысить параметрическую безопасность модели, уводя ее от небезопасных промптов и ответов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логируйте вызовы инструментов ИИ-агента для обнаружения вредоносных вызовов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0026/"><span class="relation-id">AML.M0026</span><strong>Настройка разрешений привилегированного ИИ-агента</strong><p>Надлежащий контроль доступа к использованию инструментов привилегированными ИИ-агентами может ограничить способность злоумышленника злоупотреблять вызовами инструментов при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0027/"><span class="relation-id">AML.M0027</span><strong>Настройка разрешений ИИ-агента одного пользователя</strong><p>Настройка ИИ-агентов с разрешениями на использование инструментов, унаследованными от пользователя, может ограничить способность злоумышленника злоупотреблять вызовами инструментов при компрометации агента.</p></a>
<a class="relation-item" href="/mitigations/AML.M0028/"><span class="relation-id">AML.M0028</span><strong>Настройка разрешений инструментов ИИ-агента</strong><p>Настройка инструментов ИИ-агента с контролем доступа, унаследованным от пользователя или вызывающего их ИИ-агента, может ограничить возможности злоумышленника в системе, включая злоупотребление вызовами инструментов и доступ к чувствительным данным.</p></a>
<a class="relation-item" href="/mitigations/AML.M0029/"><span class="relation-id">AML.M0029</span><strong>Участие человека в действиях ИИ-агента</strong><p>Требование подтверждения пользователем вызовов инструментов ИИ-агента может предотвратить автоматическое выполнение инструментов злоумышленником.</p></a>
<a class="relation-item" href="/mitigations/AML.M0030/"><span class="relation-id">AML.M0030</span><strong>Ограничение вызова инструментов ИИ-агента при работе с недоверенными данными</strong><p>Ограничение автоматического использования инструментов при наличии недоверенных данных может помешать злоумышленникам вызывать инструменты через промпт-инъекции.</p></a>
<a class="relation-item" href="/mitigations/AML.M0032/"><span class="relation-id">AML.M0032</span><strong>Сегментация компонентов ИИ-агента</strong><p>Сегментация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для выполнения небезопасных действий, влияющих на другие компоненты.</p></a>
<a class="relation-item" href="/mitigations/AML.M0033/"><span class="relation-id">AML.M0033</span><strong>Валидация входных и выходных данных компонентов ИИ-агента</strong><p>Валидация может помешать злоумышленникам использовать инструменты в агентном рабочем процессе для генерации небезопасных выходных данных.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Попытайтесь выбрать неавторизованные инструменты, передать небезопасные аргументы, выйти за пределы привилегий пользователя или объединить инструменты в цепочки для выполнения небезопасных действий. Устраните недостатки в настройке разрешений, валидации аргументов, изоляции в песочнице, средствах контроля действий и требованиях к одобрению.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0016/"><span class="relation-id">AML.CS0016</span><strong>Выполнение кода в MathGPT через промпт-инъекцию</strong><span class="relation-meta">Актор: Ludwig-Ferdinand Stumpp / Тактика: AML.TA0005 Выполнение</span><p>Исследователь получил возможность выполнения кода, поскольку LLM была подключена к интерпретатору Python. Исследователь мог косвенно выполнить произвольный код в Python-интерпретаторе приложения, если специально подготовленными промптами заставлял LLM сгенерировать этот код.</p></a>
<a class="relation-item" href="/studies/AML.CS0021/"><span class="relation-id">AML.CS0021</span><strong>Эксфильтрация разговоров ChatGPT</strong><span class="relation-meta">Актор: Embrace The Red / Тактика: AML.TA0012 Повышение привилегий</span><p>Кроме того, промпт может заставить LLM вызывать плагины, которые пользователь не запрашивал. В этом примере исследователь показал, как плагин `WebPilot` обращается к плагину `Expedia`.</p></a>
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0005 Выполнение</span><p>Исследователи отправляют на адрес в целевой почтовой системе письмо с состязательным самореплицирующимся промптом, или «ИИ-червем». GenAI-почтовый ассистент автоматически обрабатывает это письмо в рамках обычной генерации предлагаемого ответа. Письмо сохраняется в базе данных для RAG (генерации, дополненной извлечением), что компрометирует RAG-систему.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0012 Повышение привилегий</span><p>Исследователи скомпрометировали плагин `search_enterprise`, заставив LLM переопределить часть поведения и использовать в ответе только извлеченный объект `EmailMessage`.</p></a>
<a class="relation-item" href="/studies/AML.CS0038/"><span class="relation-id">AML.CS0038</span><strong>Внедрение инструкций для отложенного автоматического вызова инструмента ИИ-агента</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0012 Повышение привилегий</span><p>При следующем взаимодействии жертвы с Gemini вызывалось расширение Workspace.</p></a>
<a class="relation-item" href="/studies/AML.CS0039/"><span class="relation-id">AML.CS0039</span><strong>Living Off AI: промпт-инъекция через Jira Service Management</strong><span class="relation-meta">Актор: Cato CTRL / Тактика: AML.TA0012 Повышение привилегий</span><p>Вредоносный промпт запрашивал информацию, доступную ИИ-агенту через инструменты Atlassian MCP. Это приводило к вызову этих инструментов через MCP и повышало привилегии исследователей в JSM-экземпляре жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0045/"><span class="relation-id">AML.CS0045</span><strong>Эксфильтрация данных через MCP-сервер, используемый Cursor</strong><span class="relation-meta">Актор: Backslash Security Research Team / Тактика: AML.TA0012 Повышение привилегий</span><p>Промпт-инъекция задействовала возможность Cursor вызывать команды через инструмент `run_terminal_cmd`. Перед выполнением команды оболочки Cursor запросил подтверждение пользователя, что потенциально снижало риск атаки.</p></a>
<a class="relation-item" href="/studies/AML.CS0046/"><span class="relation-id">AML.CS0046</span><strong>Уничтожение данных через косвенную промпт-инъекцию, нацеленную на Claude Computer Use</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0005 Выполнение</span><p>Claude Computer Use вызвал свой инструмент `bash`, чтобы выполнить вредоносную команду.</p></a>
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учетные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0012 Повышение привилегий</span><p>Исследователь отправил ClawdBot промпт `root`; в ответ ClawdBot вызвал навык `bash`, запущенный от имени пользователя root.</p></a>
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0012 Повышение привилегий</span><p>Claude Code выполнил команду оболочки через свой инструмент `bash`.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0005 Выполнение</span><p>Промпт-инъекция заставила OpenClaw вызвать навык `bash`, чтобы загрузить и выполнить вредоносный скрипт.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0012 Повышение привилегий</span><p>Промпты исследователей вызывали инструменты ИИ-агента и были нацелены на цепочки вызовов, которые могут привести к выполнению кода.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через удаленный отравленный MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0005 Выполнение</span><p>Промпт вызывал инструмент агента, способный читать файлы из файловой системы жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0055/"><span class="relation-id">AML.CS0055</span><strong>AI ClickFix: захват управления computer-use-агентами с помощью ClickFix</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0012 Повышение привилегий</span><p>Нажатие кнопки &#34;see instructions&#34; выполняло JavaScript, который помещал вредоносную команду в буфер обмена агента. Затем агент следовал инструкциям: открывал терминал, вставлял содержимое буфера обмена и нажимал Return, выполняя команду.</p></a>
<a class="relation-item" href="/studies/AML.CS0062/"><span class="relation-id">AML.CS0062</span><strong>RCE-уязвимость в Semantic Kernel Search Plugin</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0012 Повышение привилегий</span><p>Агент Semantic Kernel вызвал инструмент поиска с вредоносным аргументом, предназначенным для выхода за пределы строки фильтра.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0012 Повышение привилегий</span><p>Вредоносный промпт заставил Gemini использовать права доступа жертвы к Calendar для вызова инструментов Google Calendar и изменения данных Calendar.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0012 Повышение привилегий</span><p>Вредоносный промпт заставил Gemini задействовать Google Home через авторизованное подключение жертвы для управления подключёнными окнами, бойлером или освещением.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0012 Повышение привилегий</span><p>Вредоносный промпт заставил Gemini использовать Android Utilities, чтобы открыть подконтрольный злоумышленнику URL в браузере жертвы и инициировать скачивание.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0012 Повышение привилегий</span><p>Вредоносный промпт заставил Gemini использовать Android Utilities и ссылку на приложение или цепочку перенаправлений, чтобы запустить Zoom.</p></a>
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Отравленные шаблоны GGUF: атака на цепочку поставок во время инференса</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0005 Выполнение</span><p>Когда отравленная модель работает в составе ИИ-агента, внедрённая инструкция меняет выбор инструментов агентом, аргументы их вызовов или порядок выполнения операций в интересах достижения цели злоумышленника, при этом позволяя продолжить выполнение легитимной задачи.</p></a>
<a class="relation-item" href="/studies/AML.CS0066/"><span class="relation-id">AML.CS0066</span><strong>ZombieAgent: атака на ChatGPT с эксфильтрацией данных</strong><span class="relation-meta">Актор: Radware Security Researchers / Тактика: AML.TA0012 Повышение привилегий</span><p>Вредоносные инструкции заставили ChatGPT задействовать возможности коннекторов и веб-доступа, доступные в рамках полномочий жертвы. В результате выполнение инструкций из промпта обеспечило доступ к информации и возможность выполнять действия, которые не были напрямую доступны исследователям.</p></a>
<a class="relation-item" href="/studies/AML.CS0066/"><span class="relation-id">AML.CS0066</span><strong>ZombieAgent: атака на ChatGPT с эксфильтрацией данных</strong><span class="relation-meta">Актор: Radware Security Researchers / Тактика: AML.TA0015 Латеральное перемещение</span><p>Вредоносные инструкции заставили ChatGPT воспроизводить промпт в новых электронных письмах или документах и распространять эти материалы среди собранных контактов. Если отравленное содержимое обрабатывал другой ИИ-агент, атака могла распространяться между пользователями или подключёнными ИИ-системами.</p></a>
<a class="relation-item" href="/studies/AML.CS0067/"><span class="relation-id">AML.CS0067</span><strong>Раскрытие секретов через Claude Code GitHub Action</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0005 Выполнение</span><p>Claude invoked its built-in Read tool on `/proc/self/environ`. Read did not execute within the Bubblewrap and scrubbed-environment boundary applied to Bash subprocesses.</p></a>
</div>
