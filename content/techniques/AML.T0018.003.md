---
atlas_id: AML.T0018.003
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Злоумышленники могут изменять шаблоны, разделители ролей, встроенные системные инструкции, настройки токенизатора, форматирование вызовов инструментов или иную включенную в состав артефакта логику, которая формирует...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 0
modified_date: "2026-07-31"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Modify Prompt Construction Logic
subtechnique_count: 0
subtechnique_of: AML.T0018
tactics:
    - AML.TA0001
    - AML.TA0006
title: Изменение логики формирования промпта
url: /techniques/AML.T0018.003/
---

Злоумышленники могут изменять шаблоны, разделители ролей, встроенные системные инструкции, настройки токенизатора, форматирование вызовов инструментов или иную включенную в состав артефакта логику, которая формирует контекст, передаваемый ИИ-модели. В файлах моделей таких форматов, как GGUF, эта логика может быть упакована вместе с весами модели в единый распространяемый артефакт. Совместимая среда выполнения инференса может интерпретировать измененную логику при обработке последующих запросов на инференс, что открывает возможность для устойчивого скрытого внедрения инструкций, изменения их приоритета, перенаправления вызовов инструментов или манипулирования выходными данными модели — без изменения ее весов.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0018/"><span class="relation-id">AML.T0018</span><strong>Манипуляция ИИ-моделью</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0018.002/"><span class="relation-id">AML.T0018.002</span><strong>Встраивание вредоносного ПО</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Отравленные шаблоны GGUF: атака на цепочку поставок во время инференса</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Злоумышленник изменяет шаблон чата, включённый в состав артефакта модели. При наличии триггера изменённый шаблон внедряет подконтрольные злоумышленнику инструкции в контекст модели, не изменяя её весов.</p></a>
</div>
