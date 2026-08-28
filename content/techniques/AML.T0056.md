---
atlas_id: AML.T0056
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2023-10-25"
description: Злоумышленники могут пытаться извлечь системный промпт большой языковой модели (LLM). Это можно сделать с помощью промпт-инъекции, чтобы заставить модель раскрыть собственный системный промпт, либо путем извлечения из...
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 4
modified_date: "2026-05-27"
platforms:
    - Generative AI
    - Agentic AI
procedure_count: 0
source_name: Extract LLM System Prompt
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0010
title: Извлечение системного промпта LLM
url: /techniques/AML.T0056/
---

Злоумышленники могут пытаться извлечь системный промпт большой языковой модели (LLM). Это можно сделать с помощью промпт-инъекции, чтобы заставить модель раскрыть собственный системный промпт, либо путем извлечения из конфигурационного файла.

Системные промпты могут быть частью конкурентного преимущества поставщика ИИ и поэтому представлять собой ценную интеллектуальную собственность, на которую могут нацеливаться злоумышленники.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0010/"><span class="relation-id">AML.TA0010</span><strong>Эксфильтрация</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0020/"><span class="relation-id">AML.M0020</span><strong>Гардрейлы для генеративного ИИ</strong><p>Гардрейлы могут предотвращать вредоносные входные данные, способные привести к извлечению метапромпта.</p></a>
<a class="relation-item" href="/mitigations/AML.M0021/"><span class="relation-id">AML.M0021</span><strong>Правила и инструкции для генеративного ИИ</strong><p>Инструкции для модели могут предписывать ей отказываться отвечать на небезопасные входные данные.</p></a>
<a class="relation-item" href="/mitigations/AML.M0022/"><span class="relation-id">AML.M0022</span><strong>Выравнивание модели генеративного ИИ</strong><p>Выравнивание модели может повысить параметрическую безопасность модели, уводя ее от небезопасных промптов и ответов.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Тестируйте интерфейсы модели и приложения на возможность раскрытия системных инструкций, политик, определений инструментов или скрытого контекста. Удаляйте встроенные секреты и совершенствуйте изоляцию конфигурации и средства контроля выходных данных.</p></a>
</div>
