---
atlas_id: AML.T0051.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2023-10-25"
description: Злоумышленник может внедрять промпты напрямую, действуя как пользователь LLM. Такой тип инъекции может использоваться злоумышленником для закрепления в системе или для злоупотребления самой LLM, например для генерации...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 12
source_name: Direct
subtechnique_count: 0
subtechnique_of: AML.T0051
tactics:
    - AML.TA0005
title: Прямая промпт-инъекция
url: /techniques/AML.T0051.000/
---

Злоумышленник может внедрять промпты напрямую, действуя как пользователь LLM. Такой тип инъекции может использоваться злоумышленником для закрепления в системе или для злоупотребления самой LLM, например для генерации вредоносного содержимого.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0005/"><span class="relation-id">AML.TA0005</span><strong>Выполнение</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0051/"><span class="relation-id">AML.T0051</span><strong>Промпт-инъекция в LLM</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0051.001/"><span class="relation-id">AML.T0051.001</span><strong>Косвенная промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0051.002/"><span class="relation-id">AML.T0051.002</span><strong>Триггерная промпт-инъекция</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0024/"><span class="relation-id">AML.M0024</span><strong>Логирование телеметрии ИИ</strong><p>Логирование телеметрии может помочь выявить отправку небезопасных промптов в LLM.</p></a>
<a class="relation-item" href="/mitigations/AML.M0033/"><span class="relation-id">AML.M0033</span><strong>Валидация входных и выходных данных компонентов ИИ-агента</strong><p>Валидация может помешать злоумышленникам выполнять промпт-инъекции, способные повлиять на агентные рабочие процессы.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0016/"><span class="relation-id">AML.CS0016</span><strong>Выполнение кода в MathGPT через промпт-инъекцию</strong><span class="relation-meta">Актор: Ludwig-Ferdinand Stumpp / Тактика: AML.TA0005 Выполнение</span><p>Исследователь вручную составлял состязательные промпты, чтобы проверить, уязвима ли модель к промпт-инъекции и действительно ли приложение напрямую выполняет код, сгенерированный GPT-3.</p></a>
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0005 Выполнение</span><p>Исследователи тестируют промпты через публичные API моделей, чтобы найти работающие промпт-инъекции.</p></a>
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0005 Выполнение</span><p>Злоумышленник отправлял вредоносные промпты напрямую в LLM в рамках любой текущей переписки жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0041/"><span class="relation-id">AML.CS0041</span><strong>Бэкдор в файле правил: атака на цепочку поставки ИИ-ассистентов для программирования</strong><span class="relation-meta">Актор: Pillar Security / Тактика: AML.TA0005 Выполнение</span><p>При следующей инициализации ИИ-ассистента для программирования файл правил считывался, и вредоносный промпт выполнялся.</p></a>
<a class="relation-item" href="/studies/AML.CS0043/"><span class="relation-id">AML.CS0043</span><strong>Прототип вредоносного ПО со встроенной промпт-инъекцией</strong><span class="relation-meta">Актор: Unknown Threat Actor / Тактика: AML.TA0005 Выполнение</span><p>Когда LLM-инструмент обнаружения или анализа вредоносного ПО обрабатывает бинарный файл Skynet, срабатывает встроенный промпт.</p></a>
<a class="relation-item" href="/studies/AML.CS0047/"><span class="relation-id">AML.CS0047</span><strong>Код для развертывания деструктивного ИИ-агента обнаружен в расширении Amazon Q для VS Code</strong><span class="relation-meta">Актор: lkmanka58 (GitHub user) / Тактика: AML.TA0005 Выполнение</span><p>Агент Amazon Q был развернут с промпт-инъекцией, которая предписывала ему выполнить деструктивные действия в файловой системе и облачной среде жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0049/"><span class="relation-id">AML.CS0049</span><strong>Компрометация цепочки поставки через отравленный навык ClawdBot</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0005 Выполнение</span><p>Claude Code прочитал все файлы навыка и выполнил вредоносный промпт из `rules/logic.md`.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0005 Выполнение</span><p>Когда жертва взаимодействовала с OpenClaw, выполнялся измененный системный промпт с инструкциями исследователей.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0005 Выполнение</span><p>Исследователи напрямую отправляли ИИ-агенту вредоносные инструкции.</p></a>
<a class="relation-item" href="/studies/AML.CS0054/"><span class="relation-id">AML.CS0054</span><strong>Эксфильтрация данных через удаленный отравленный MCP-инструмент</strong><span class="relation-meta">Актор: Invariant Labs / Тактика: AML.TA0005 Выполнение</span><p>Когда пользователь вызывал удаленный MCP-инструмент, промпт-инъекция, скрытая в docstring-описании, выполнялась локально.</p></a>
<a class="relation-item" href="/studies/AML.CS0060/"><span class="relation-id">AML.CS0060</span><strong>Межсайтовый скриптинг (XSS) через манипуляцию промптом в ИИ-чат-боте Lenovo</strong><span class="relation-meta">Актор: Cybernews Research Team / Тактика: AML.TA0005 Выполнение</span><p>Lena выполнила подконтрольные злоумышленнику инструкции форматирования и сгенерировала HTML-ответ, содержащий вредоносную нагрузку. Ответ сохранился в ветке чата.</p></a>
<a class="relation-item" href="/studies/AML.CS0062/"><span class="relation-id">AML.CS0062</span><strong>RCE-уязвимость в Semantic Kernel Search Plugin</strong><span class="relation-meta">Актор: Microsoft Defender Security Research Team / Тактика: AML.TA0005 Выполнение</span><p>Исследователи отправили агенту подготовленный промпт. Промпт-инъекция заставила модель подготовить вызов инструмента поиска с вредоносным аргументом.</p></a>
</div>
