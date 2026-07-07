---
atlas_id: AML.T0019
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут отравить обучающие данные и опубликовать их в публичном месте. Отравленный набор данных может быть новым набором данных или отравленным вариантом существующего набора данных с открытым исходным...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 3
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 1
source_name: Publish Poisoned Datasets
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Публикация отравленных наборов данных
url: /techniques/AML.T0019/
---

Злоумышленники могут [отравить обучающие данные](/techniques/AML.T0020) и опубликовать их в публичном месте.

Отравленный набор данных может быть новым набором данных или отравленным вариантом существующего набора данных с открытым исходным кодом.

Эти данные могут попасть в систему организации-жертвы через [компрометацию цепочки поставок ИИ](/techniques/AML.T0010).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0014/"><span class="relation-id">AML.M0014</span><strong>Проверка ИИ-артефактов</strong><p>Проверяйте достоверность опубликованных данных, чтобы не использовать отравленные данные, создающие уязвимости.</p></a>
<a class="relation-item" href="/mitigations/AML.M0023/"><span class="relation-id">AML.M0023</span><strong>Ведомость материалов ИИ</strong><p>AI BOM может помочь пользователям выявлять недоверенные артефакты моделей.</p></a>
<a class="relation-item" href="/mitigations/AML.M0025/"><span class="relation-id">AML.M0025</span><strong>Поддержание происхождения наборов данных ИИ</strong><p>Ведение подробной истории наборов данных может помочь выявлять использование отравленных наборов данных из публичных источников.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0025/"><span class="relation-id">AML.CS0025</span><strong>Отравление крупномасштабных веб-датасетов: атака split-view</strong><span class="relation-meta">Актор: Researchers from Google Deepmind, ETH Zurich, NVIDIA, Robust Intelligence, and Google / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Затем злоумышленник может разместить отравленные данные на подконтрольных доменах. В этом упражнении исследователи отслеживали обращения к подконтрольным URL, чтобы показать, что датасет действительно продолжает использоваться.</p></a>
</div>
