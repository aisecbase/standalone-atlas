---
atlas_id: AML.T0050
atlas_type: technique
attack_ref_id: T1059
attack_ref_url: https://attack.mitre.org/techniques/T1059/
created_date: "2023-02-28"
description: Злоумышленники могут злоупотреблять командными и скриптовыми интерпретаторами для выполнения команд, сценариев или бинарных файлов. Эти интерфейсы и языки предоставляют способы взаимодействия с компьютерными системами...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 7
source_name: Command and Scripting Interpreter
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0005
title: Интерпретатор команд и сценариев
url: /techniques/AML.T0050/
---

Злоумышленники могут злоупотреблять командными и скриптовыми интерпретаторами для выполнения команд, сценариев или бинарных файлов. Эти интерфейсы и языки предоставляют способы взаимодействия с компьютерными системами и широко встречаются на разных платформах. Большинство систем имеют встроенный интерфейс командной строки и возможности выполнения сценариев: например, дистрибутивы macOS и Linux включают один из вариантов Unix Shell, а установки Windows — Windows Command Shell и PowerShell.

Существуют также кроссплатформенные интерпретаторы, например Python, а также интерпретаторы, обычно связанные с клиентскими приложениями, такие как JavaScript и Visual Basic.

Злоумышленники могут использовать эти технологии различными способами как средство выполнения произвольных команд. Команды и сценарии могут быть встроены в полезные нагрузки для первичного доступа, доставляемые жертвам в виде документов-приманок, или во вторичные полезные нагрузки, загружаемые из существующей C2-инфраструктуры. Злоумышленники также могут выполнять команды через интерактивные терминалы или оболочки, а также использовать различные удаленные сервисы для удаленного выполнения.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0011/"><span class="relation-id">AML.M0011</span><strong>Ограничение загрузки библиотек</strong><p>Ограничивайте загрузку библиотек так, чтобы интерпретаторы команд и скриптов не могли загружать недоверенные библиотеки для выполнения кода.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0050/"><span class="relation-id">AML.CS0050</span><strong>Удаленное выполнение кода (RCE) в OpenClaw в один клик</strong><span class="relation-meta">Актор: DepthFirst / Тактика: AML.TA0005 Выполнение</span><p>Вредоносный скрипт добивался удаленного выполнения кода, отправляя запрос `node.invoke` (RPC-механизм OpenClaw) к API OpenClaw.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0005 Выполнение</span><p>Код из промптов исследователей выполнялся в изолированном Python-интерпретаторе.</p></a>
<a class="relation-item" href="/studies/AML.CS0060/"><span class="relation-id">AML.CS0060</span><strong>Межсайтовый скриптинг (XSS) через манипуляцию промптом в ИИ-чат-боте Lenovo</strong><span class="relation-meta">Актор: Cybernews Research Team / Тактика: AML.TA0005 Выполнение</span><p>Сохраненный HTML включал исполняемый в браузере JavaScript, который запускался в браузере сотрудника поддержки при отображении стенограммы.</p></a>
<a class="relation-item" href="/studies/AML.CS0061/"><span class="relation-id">AML.CS0061</span><strong>AI in the Middle: веб-сервисы ИИ как ретрансляторы C2</strong><span class="relation-meta">Актор: Check Point Research / Тактика: AML.TA0005 Выполнение</span><p>ИИ-сервис кратко излагал ответ подконтрольного злоумышленнику сайта, а имплант выполнял извлеченные команды. В демонстрационном примере команда запускала Calculator с помощью `cmd.exe /c calc.exe`; реальный имплант мог бы выполнять другие команды, загружать полезные нагрузки, переходить в режим ожидания или собирать дополнительные данные.</p></a>
<a class="relation-item" href="/studies/AML.CS0062/"><span class="relation-id">AML.CS0062</span><strong>RCE-уязвимость в Semantic Kernel Search Plugin</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0005 Выполнение</span><p>Значение фильтра обрабатывалось как lambda-выражение Python. Из-за вредоносного форматирования в подконтрольном злоумышленнику аргументе эта обработка становилась приемником инъекции: ввод исследователей выходил за рамки предусмотренной логики сравнения и приводил к удаленному выполнению кода.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0005 Выполнение</span><p>Через доступный из интернета стенд проверки кода агенты выполняли shell-команды и передавали код на C и Python.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0005 Выполнение</span><p>Агенты задействовали путь выполнения через Jinja2 для запуска в продакшен-поде Dataset Server заранее размещённых команд на Python и shell-команд. Полученный вывод они использовали для корректировки последующих команд.</p></a>
</div>
