---
atlas_id: AML.T0069.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут выявлять ключевые слова, имеющие особое значение для большой языковой модели (LLM), например имена функций или объектов. Позднее они могут использоваться, чтобы запутать LLM или манипулировать ею,...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 2
source_name: System Instruction Keywords
subtechnique_count: 0
subtechnique_of: AML.T0069
tactics:
    - AML.TA0008
title: Ключевые слова системных инструкций
url: /techniques/AML.T0069.001/
---

Злоумышленники могут выявлять ключевые слова, имеющие особое значение для большой языковой модели (LLM), например имена функций или объектов. Позднее они могут использоваться, чтобы запутать LLM или манипулировать ею, вызывая некорректное поведение и обращения к плагинам, к которым у LLM есть доступ.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0008/"><span class="relation-id">AML.TA0008</span><strong>Выявление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0069/"><span class="relation-id">AML.T0069</span><strong>Выявление системной информации LLM</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0069.000/"><span class="relation-id">AML.T0069.000</span><strong>Наборы специальных символов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0069.002/"><span class="relation-id">AML.T0069.002</span><strong>Системный промпт</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Restrict access to stored system instructions and tool definitions containing privileged keywords.</p></a>
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Гардрейлы для генеративного ИИ</strong><p>Фильтруйте ответы, раскрывающие системные ключевые слова, имена инструментов, определения функций или скрытые инструкции.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0008 Выявление</span><p>Исследуя ответы Copilot, исследователи выявили плагины и конкретные функции, к которым Copilot имеет доступ. Среди них были функция `search_enterprise` и объект `EmailMessage`.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0008 Выявление</span><p>Исследователи обнаружили конкретные управляющие последовательности OpenClawd, включая `&lt;&lt;&lt;end_tool_call_result&gt;&gt;&gt;`, `&lt;&lt;&lt;start_user_message&gt;&gt;&gt;`, `&lt;&lt;&lt;end_user_message&gt;&gt;&gt;`, `&lt;think&gt;` и `&lt;/think&gt;`.</p></a>
</div>
