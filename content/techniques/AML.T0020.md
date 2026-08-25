---
atlas_id: AML.T0020
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут пытаться отравить наборы данных, используемые ИИ-моделью, изменяя исходные данные или их метки. Это позволяет злоумышленнику встроить в ИИ-модели, обученные на этих данных, уязвимости, которые...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 7
modified_date: "2026-07-31"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: Training Data Poisoning
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0006
title: Отравление обучающих данных
url: /techniques/AML.T0020/
---

Злоумышленники могут пытаться отравить наборы данных, используемые ИИ-моделью, изменяя исходные данные или их метки.

Это позволяет злоумышленнику встроить в ИИ-модели, обученные на этих данных, уязвимости, которые может быть трудно обнаружить.

Атаки отравления данных могут требовать или не требовать изменения меток.

Встроенная уязвимость активируется позднее образцами данных с [бэкдор-триггером](/techniques/AML.T0043.004).

Отравленные данные могут быть внедрены через [компрометацию цепочки поставок ИИ](/techniques/AML.T0010), либо данные могут быть отравлены после того, как злоумышленник получит [первичный доступ](/tactics/AML.TA0004) к системе.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0006/"><span class="relation-id">AML.TA0006</span><strong>Закрепление</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0001/"><span class="relation-id">AML.M0001</span><strong>Ограничение публикации артефактов модели</strong><p>Опубликованные наборы данных могут стать целью атак отравления.</p></a>
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа может предотвращать подмену ML-артефактов и несанкционированное копирование.</p></a>
<a class="relation-item" href="/mitigations/AML.M0007/"><span class="relation-id">AML.M0007</span><strong>Санитизация обучающих данных</strong><p>Выявляйте изменения данных и меток, которые могут вызвать состязательный дрейф модели или бэкдор-атаки.</p></a>
<a class="relation-item" href="/mitigations/AML.M0008/"><span class="relation-id">AML.M0008</span><strong>Валидация ИИ-модели</strong><p>Тщательная оценка ИИ-модели может повысить уверенность в том, что модель не была отравлена.</p></a>
<a class="relation-item" href="/mitigations/AML.M0023/"><span class="relation-id">AML.M0023</span><strong>Ведомость материалов ИИ</strong><p>AI BOM может помочь пользователям выявлять недоверенные артефакты моделей.</p></a>
<a class="relation-item" href="/mitigations/AML.M0025/"><span class="relation-id">AML.M0025</span><strong>Поддержание происхождения наборов данных ИИ</strong><p>Сведения о происхождении наборов данных могут защищать от отравления обучающих данных.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>В контролируемых условиях внедрите отравленные записи или триггеры в репрезентативные конвейеры данных. Проверьте и усовершенствуйте отслеживание происхождения данных, их санитизацию и проверку, обнаружение дрейфа, валидацию модели и механизмы возврата к предыдущему состоянию.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0002/"><span class="relation-id">AML.CS0002</span><strong>Отравление VirusTotal</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0006 Закрепление</span><p>Несколько вендоров начали классифицировать файлы как относящиеся к этому семейству программ-вымогателей, хотя большинство из них не запускались. Мутированные образцы отравили набор данных, который модели машинного обучения используют для выявления и классификации этого семейства программ-вымогателей.</p></a>
<a class="relation-item" href="/studies/AML.CS0009/"><span class="relation-id">AML.CS0009</span><strong>Отравление Tay</strong><span class="relation-meta">Актор: 4chan Users / Тактика: AML.TA0006 Закрепление</span><p>Многократно взаимодействуя с Tay с использованием расистской и оскорбительной лексики, злоумышленники смогли сместить набор данных Tay в сторону такой же лексики. Для этого они использовали функцию &#34;repeat after me&#34; — команду, которая заставляла Tay повторять все, что ей говорили.</p></a>
<a class="relation-item" href="/studies/AML.CS0025/"><span class="relation-id">AML.CS0025</span><strong>Отравление крупномасштабных веб-датасетов: атака split-view</strong><span class="relation-meta">Актор: Researchers from Google Deepmind, ETH Zurich, NVIDIA, Robust Intelligence, and Google / Тактика: AML.TA0006 Закрепление</span><p>An adversary could create poisoned training data to replace expired portions of the dataset.</p></a>
</div>
