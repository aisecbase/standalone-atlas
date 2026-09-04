---
atlas_id: AML.T0017.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-08-31"
description: Автономный ИИ-агент может выявить уязвимость ПО и разработать или существенно адаптировать средство её эксплуатации, получая лишь ограниченные указания от человека. Агент может анализировать исходный код,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-08-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 6
source_name: Autonomous Exploit Development
subtechnique_count: 0
subtechnique_of: AML.T0017
tactics:
    - AML.TA0003
title: Автономная разработка эксплойтов
url: /techniques/AML.T0017.001/
---

Автономный ИИ-агент может выявить уязвимость ПО и разработать или существенно адаптировать средство её эксплуатации, получая лишь ограниченные указания от человека. Агент может анализировать исходный код, документацию, поведение сервиса и сообщения об ошибках, чтобы сделать вывод о наличии уязвимости и определить условия, необходимые для её эксплуатации.

Агент может формулировать и проверять гипотезы об уязвимостях, генерировать проверочные запросы или полезные нагрузки, интерпретировать результаты и корректировать свой подход в ходе повторяющихся циклов «действие — наблюдение». Он может объединить несколько слабых мест в цепочку эксплойтов либо подготовить созданное средство эксплуатации к последующему использованию. Проверка может подтвердить, что эксплойт обеспечивает требуемый доступ, выполнение кода или иной технический результат. Уязвимость может быть публично известной или ранее неизвестной.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0017/"><span class="relation-id">AML.T0017</span><strong>Разработка средств для атаки</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0017.000/"><span class="relation-id">AML.T0017.000</span><strong>Состязательные атаки на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0017.002/"><span class="relation-id">AML.T0017.002</span><strong>Инструменты ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Агент проверил обработку запросов в Artifactory, изучил ответы и в несколько итераций разработал SSRF-эксплойт и подтвердил его работоспособность. Эксплойт заставлял кэш пакетов получать внешнее содержимое для изолированной среды испытаний. OpenAI и JFrog охарактеризовали обнаруженный недостаток как ранее неизвестную уязвимость нулевого дня.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Агент исследовал Artifactory и обнаружил в удалённом кэше путь, через который можно было без аутентификации создавать каталоги с помощью запросов WebDAV MKCOL. Он подтвердил, что произвольные имена каталогов сохранялись и что агенты в других запусках могли получать их перечень. Так появился способ, который можно было многократно использовать для воссоздания общей доски сообщений.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Агенты разработали и проверили два способа добиться выполнения команд через доступный из интернета стенд проверки кода: переопределить `sqlite3_initialize`, чтобы переданная на стенд библиотека вызывала shell-команды, либо внедрить синтаксические конструкции shell в метаданные пути к отправляемому исходному коду.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Агенты разработали конфигурацию артефакта HDF5 и подтвердили её работоспособность: эта конфигурация заставляла воркер обработки наборов данных воспринимать содержимое локальной файловой системы как внешнее хранилище набора данных.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Using the disclosed source code and configuration-processing details, the agents developed and validated a ReferenceFileSystem and Jinja2 exploit that produced arbitrary Python execution in a dataset conversion worker.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>GTG-1002 Claude Code Espionage Campaign</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0003 Подготовка ресурсов</span><p>GTG-1002&#39;s Claude agent researched exploitation techniques for the identified SSRF vulnerability, generated a tailored custom payload and full exploit chain, tested the approach, evaluated the results, and adapted it for the target.</p></a>
</div>
