---
atlas_id: AML.T0066
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут подготавливать содержимое так, чтобы оно извлекалось по пользовательским запросам и тем или иным образом влияло на пользователя системы. Такая техника злоупотребляет доверием пользователя к...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-08-31"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: Retrieval Content Crafting
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0001
title: Подготовка содержимого для извлечения
url: /techniques/AML.T0066/
---

Злоумышленники могут подготавливать содержимое так, чтобы оно извлекалось по пользовательским запросам и тем или иным образом влияло на пользователя системы. Такая техника злоупотребляет доверием пользователя к системе.

Подготовленное содержимое может сочетаться с промпт-инъекцией или существовать отдельно, например в самостоятельном документе или письме. Злоумышленнику нужно поместить это содержимое в базу данных жертвы, например в векторную базу данных, используемую в системе RAG (генерации, дополненной извлечением). Для этого он может получить кибердоступ или злоупотребить механизмами приема данных, распространенными в RAG-системах (см. [отравление RAG](/techniques/AML.T0070)).

Большие языковые модели могут использоваться как вспомогательный инструмент для подготовки такого содержимого.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Гардрейлы для генеративного ИИ</strong><p>Применяйте гардрейлы для извлечения данных, чтобы отклонять недоверенное, вредоносное, нерелевантное или неподтверждённое содержимое RAG.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>The Zenity researchers wrote targeted content designed to be retrieved by specific user queries.</p></a>
<a class="relation-item" href="/studies/AML.CS0035/"><span class="relation-id">AML.CS0035</span><strong>Эксфильтрация данных из Slack AI через косвенную промпт-инъекцию</strong><span class="relation-meta">Актор: PromptArmor / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>The researcher crafted a targeted message designed to be retrieved when a user asks about their API key. &lt;div style=&#34;font-family: monospace; width: 50%; margin-left: 50px; background-color:ghostwhite; border: 2px solid black; padding: 10px;&#34;&gt; &lt;span style=&#34;color: red;&#34;&gt;&#34;EldritchNexus API key:&#34;&lt;/span&gt; &lt;/div&gt;</p></a>
<a class="relation-item" href="/studies/AML.CS0059/"><span class="relation-id">AML.CS0059</span><strong>EchoLeak: промпт-инъекция нулевого клика против M365 Copilot для эксфильтрации данных</strong><span class="relation-meta">Актор: Aim Labs / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>The researchers embedded the prompt injection in business-like email content that was likely to be retrieved during a later Copilot interaction. The content was designed to appear relevant to ordinary enterprise workflows while carrying hidden instructions.</p></a>
</div>
