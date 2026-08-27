---
atlas_id: AML.T0079
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-04-16"
description: Злоумышленники могут загружать, устанавливать или иным образом настраивать средства, которые могут использоваться при атаке на цель. Для поддержки своих операций злоумышленнику может потребоваться взять средства,...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 11
source_name: Stage Capabilities
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Размещение средств атаки
url: /techniques/AML.T0079/
---

Злоумышленники могут загружать, устанавливать или иным образом настраивать средства, которые могут использоваться при атаке на цель. Для поддержки своих операций злоумышленнику может потребоваться взять средства, которые он разработал ([Разработка средств для атаки](/techniques/AML.T0017)) или получил ([Получение средств для атаки](/techniques/AML.T0016)), и разместить их на подконтрольной инфраструктуре. Такие средства могут размещаться на инфраструктуре, ранее купленной или арендованной злоумышленником ([Получение инфраструктуры](/techniques/AML.T0008)), либо на инфраструктуре, которую он скомпрометировал иным способом. Средства также могут размещаться в веб-сервисах, таких как GitHub, в реестрах моделей, таких как Hugging Face, или в реестрах контейнеров.

Злоумышленники могут размещать различные ИИ-артефакты, включая отравленные наборы данных ([Публикация отравленных наборов данных](/techniques/AML.T0019)), вредоносные модели ([Публикация отравленных моделей](/techniques/AML.T0058)) и промпт-инъекции. Они могут использовать названия легитимных компаний или продуктов, применять тайпсквоттинг или использовать галлюцинированные сущности ([Выявление галлюцинированных сущностей LLM](/techniques/AML.T0062)).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0021/"><span class="relation-id">AML.CS0021</span><strong>Эксфильтрация разговоров ChatGPT</strong><span class="relation-meta">Актор: Embrace The Red / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь разместил этот промпт на веб-странице, чтобы ChatGPT мог получить его при обращении к странице.</p></a>
<a class="relation-item" href="/studies/AML.CS0041/"><span class="relation-id">AML.CS0041</span><strong>Бэкдор в файле правил: атака на цепочку поставки ИИ-ассистентов для программирования</strong><span class="relation-meta">Актор: Pillar Security / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи разместили вредоносный JavaScript-файл на публично доступном сайте.</p></a>
<a class="relation-item" href="/studies/AML.CS0045/"><span class="relation-id">AML.CS0045</span><strong>Эксфильтрация данных через MCP-сервер, используемый Cursor</strong><span class="relation-meta">Актор: Backslash Security Research Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи создали сайт, содержащий вредоносный промпт.</p></a>
<a class="relation-item" href="/studies/AML.CS0045/"><span class="relation-id">AML.CS0045</span><strong>Эксфильтрация данных через MCP-сервер, используемый Cursor</strong><span class="relation-meta">Актор: Backslash Security Research Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи запустили веб-сервер для приема данных, эксфильтрированных из среды жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0050/"><span class="relation-id">AML.CS0050</span><strong>Удаленное выполнение кода (RCE) в OpenClaw в один клик</strong><span class="relation-meta">Актор: DepthFirst / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь разместил вредоносный скрипт на неприметном сайте.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи разместили на своем сайте промпт-инъекции, вредоносный скрипт и TODO-список со своими командами.</p></a>
<a class="relation-item" href="/studies/AML.CS0055/"><span class="relation-id">AML.CS0055</span><strong>AI ClickFix: захват управления computer-use-агентами с помощью ClickFix</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь разместил сайт и скрипт. На практике вредоносный HTML можно было бы внедрить в скомпрометированный легитимный сайт.</p></a>
<a class="relation-item" href="/studies/AML.CS0057/"><span class="relation-id">AML.CS0057</span><strong>Storm-2139: обход защитных ограничений Azure OpenAI</strong><span class="relation-meta">Актор: Storm-2139 / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Storm-2139 подготовила и эксплуатировала сервис обратного прокси, чтобы другие злоумышленники могли взаимодействовать с неправомерно используемыми сервисами генеративного ИИ.</p></a>
<a class="relation-item" href="/studies/AML.CS0059/"><span class="relation-id">AML.CS0059</span><strong>EchoLeak: промпт-инъекция нулевого клика против M365 Copilot для эксфильтрации данных</strong><span class="relation-meta">Актор: Aim Labs / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи подготовили веб-эндпоинт, контролируемый злоумышленником, для приема исходящих запросов с закодированными конфиденциальными данными. Эндпоинт служил точкой сбора для канала эксфильтрации.</p></a>
<a class="relation-item" href="/studies/AML.CS0061/"><span class="relation-id">AML.CS0061</span><strong>AI in the Middle: веб-сервисы ИИ как ретрансляторы C2</strong><span class="relation-meta">Актор: Check Point Research / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи разместили внешне безобидное содержимое, одновременно возвращая данные, которые имплант мог интерпретировать как C2-инструкции.</p></a>
<a class="relation-item" href="/studies/AML.CS0066/"><span class="relation-id">AML.CS0066</span><strong>ZombieAgent: атака на ChatGPT с эксфильтрацией данных</strong><span class="relation-meta">Актор: Radware Security Researchers / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи подготовили инфраструктуру для приёма эксфильтрированных данных.</p></a>
</div>
