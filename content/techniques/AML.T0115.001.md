---
atlas_id: AML.T0115.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-07-31"
description: Злоумышленники могут публиковать отравленные ИИ-модели через реестры моделей, репозитории кода или другие каналы распространения моделей. Модель может быть вновь созданной либо представлять собой модифицированный...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-07-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 5
source_name: Models
subtechnique_count: 0
subtechnique_of: AML.T0115
tactics:
    - AML.TA0003
title: Модели
url: /techniques/AML.T0115.001/
---

Злоумышленники могут публиковать отравленные ИИ-модели через реестры моделей, репозитории кода или другие каналы распространения моделей. Модель может быть вновь созданной либо представлять собой модифицированный вариант легитимной модели; её веса, конфигурации, архитектура, сериализованный код или другие компоненты могут быть изменены так, чтобы вызывать вредоносное поведение или обеспечивать выполнение вредоносного кода.

Впоследствии жертвы могут скачать и интегрировать модель в ходе [компрометации цепочки поставок ИИ](/techniques/AML.T0010).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0115/"><span class="relation-id">AML.T0115</span><strong>Публикация отравленных ИИ-артефактов</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0115.000/"><span class="relation-id">AML.T0115.000</span><strong>Наборы данных</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0115.002/"><span class="relation-id">AML.T0115.002</span><strong>Инструменты ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0008/"><span class="relation-id">AML.M0008</span><strong>Валидация ИИ-модели</strong><p>Model repositories evaluate submissions for backdoors, data leakage, adversarial influence, and unexpected behavior before listing.</p></a>
<a class="relation-item" href="/mitigations/AML.M0016/"><span class="relation-id">AML.M0016</span><strong>Сканирование уязвимостей</strong><p>Model registries scan uploaded models for unsafe serialization, embedded code, malware, and known vulnerabilities before listing.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0019/"><span class="relation-id">AML.CS0019</span><strong>PoisonGPT</strong><span class="relation-meta">Актор: Mithril Security Researchers / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The researchers uploaded the PoisonGPT model back to HuggingFace under a similar repository name as the original model, missing one letter.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The researcher re-uploaded the manipulated model to the Hugging Face repository.</p></a>
<a class="relation-item" href="/studies/AML.CS0031/"><span class="relation-id">AML.CS0031</span><strong>Вредоносные модели на Hugging Face</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0003 Подготовка ресурсов</span><p>The adversary uploaded the model to Hugging Face. In both instances observed by the ReversingLab, the malicious models did not make any attempt to mimic a popular legitimate model.</p></a>
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Отравленные шаблоны GGUF: атака на цепочку поставок во время инференса</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник публикует изменённый артефакт модели в репозитории моделей или через другой канал распространения, используемый разработчиками и организациями на последующих этапах цепочки поставок.</p></a>
<a class="relation-item" href="/studies/AML.CS0065/"><span class="relation-id">AML.CS0065</span><strong>Атака на цепочку поставок через повторное использование пространства имён модели</strong><span class="relation-meta">Актор: Unit 42 Researchers / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Специалисты Unit 42 загрузили вредоносную модель в повторно зарегистрированное пространство имён, используя исходный идентификатор Author/ModelName.</p></a>
</div>
