---
atlas_id: AML.T0061
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-03-12"
description: Злоумышленник может использовать тщательно подготовленную промпт-инъекцию в LLM, рассчитанную на то, чтобы заставить LLM воспроизвести промпт как часть своего ответа. Это позволяет промпту распространяться на другие...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 3
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: LLM Prompt Self-Replication
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0006
title: Саморепликация промпта LLM
url: /techniques/AML.T0061/
---

Злоумышленник может использовать тщательно подготовленную [промпт-инъекцию в LLM](/techniques/AML.T0051), рассчитанную на то, чтобы заставить LLM воспроизвести промпт как часть своего ответа. Это позволяет промпту распространяться на другие LLM и сохраняться в системе. Самореплицирующийся промпт обычно сочетается с другими вредоносными инструкциями, например [джейлбрейком LLM](/techniques/AML.T0054) или [утечкой данных из LLM](/techniques/AML.T0057).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Гардрейлы для генеративного ИИ</strong><p>Гардрейлы могут помогать предотвращать атаки репликации во входных и выходных данных модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0021/"><span class="relation-id">AML.M0021</span><strong>Правила и инструкции для генеративного ИИ</strong><p>Инструкции могут направлять модель к созданию более безопасных выходных данных и предотвращать генерацию самореплицирующихся выходных данных.</p></a>
<a class="relation-item" href="/mitigations/AML.M0022/"><span class="relation-id">AML.M0022</span><strong>Выравнивание модели генеративного ИИ</strong><p>Выравнивание модели может повысить защищенность моделей от атак с самореплицирующимися промптами.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0024/"><span class="relation-id">AML.CS0024</span><strong>Червь Morris II: атака на основе RAG</strong><span class="relation-meta">Актор: Stav Cohen, Ron Bitton, Ben Nassi / Тактика: AML.TA0006 Закрепление</span><p>Самореплицирующаяся часть промпта заставляет сгенерированный ответ содержать вредоносный промпт, благодаря чему червь может распространяться дальше.</p></a>
</div>
