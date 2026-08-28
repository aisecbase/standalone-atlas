---
atlas_id: AML.T0069.000
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленники могут выявлять разделители и наборы специальных символов, используемые большой языковой моделью. Например, разделители, применяемые в приложениях RAG (генерации, дополненной извлечением), чтобы отличать...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 3
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 2
source_name: Special Character Sets
subtechnique_count: 0
subtechnique_of: AML.T0069
tactics:
    - AML.TA0008
title: Наборы специальных символов
url: /techniques/AML.T0069.000/
---

Злоумышленники могут выявлять разделители и наборы специальных символов, используемые большой языковой моделью. Например, разделители, применяемые в приложениях RAG (генерации, дополненной извлечением), чтобы отличать контекст от пользовательских промптов. Позднее эти данные могут быть использованы для того, чтобы запутать большую языковую модель или манипулировать ею, вызывая некорректное поведение.


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
<a class="relation-item" href="/techniques/AML.T0069.001/"><span class="relation-id">AML.T0069.001</span><strong>Ключевые слова системных инструкций</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0069.002/"><span class="relation-id">AML.T0069.002</span><strong>Системный промпт</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Restrict access to stored prompt templates and configurations containing internal delimiters.</p></a>
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Authenticate and monitor access to production models and prompt configuration.</p></a>
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Гардрейлы для генеративного ИИ</strong><p>Фильтруйте ответы, раскрывающие скрытое форматирование промпта, разделители или внутренние инструкции.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0008 Выявление</span><p>Исследуя ответы Copilot, исследователи выявили специальные разделители и маркеры, например `**`, `**END**`, `Actual Snippet:` и `[^1^]`. Эти строки используются как служебные признаки для отделения разных частей промпта Copilot друг от друга.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0008 Выявление</span><p>Исследователи выявили специальные символы, например `&lt;&lt;&lt;` и `&gt;&gt;&gt;`, которые OpenClawd использует для обозначения управляющих последовательностей.</p></a>
</div>
