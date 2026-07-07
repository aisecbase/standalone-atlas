---
atlas_id: AML.M0013
atlas_type: mitigation
attack_ref_id: M1045
attack_ref_url: https://attack.mitre.org/mitigations/M1045/
category:
    - Technical - Cyber
created_date: "2023-04-12"
description: Обеспечивайте целостность бинарных файлов и приложений с помощью проверки цифровой подписи, чтобы предотвратить выполнение недоверенного кода. Злоумышленники могут встраивать вредоносный код в ИИ-ПО или модели....
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Deployment
modified_date: "2026-03-19"
source_name: Code Signing
technique_count: 8
title: Подписание кода
url: /mitigations/AML.M0013/
---

Обеспечивайте целостность бинарных файлов и приложений с помощью проверки цифровой подписи, чтобы предотвратить выполнение недоверенного кода.
Злоумышленники могут встраивать вредоносный код в ИИ-ПО или модели.
Разработчикам также следует криптографически подписывать компоненты SBOM и AIBOM, которые фиксируют происхождение модели или данных.
Обязательное применение подписания кода может предотвратить компрометацию цепочки поставок ИИ и выполнение вредоносного кода.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010.001/"><span class="relation-id">AML.T0010.001</span><strong>ПО для ИИ</strong><p>Требуйте корректной подписи драйверов и ML-фреймворков.</p></a>
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><p>Требуйте корректной подписи файлов модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0011.000/"><span class="relation-id">AML.T0011.000</span><strong>Небезопасные ИИ-артефакты</strong><p>Предотвращайте выполнение ML-артефактов, которые не подписаны должным образом.</p></a>
<a class="relation-item" href="/techniques/AML.T0011.001/"><span class="relation-id">AML.T0011.001</span><strong>Вредоносный пакет</strong><p>Подписание кода дает гарантию, что программный пакет не был изменен после подписания.</p></a>
<a class="relation-item" href="/techniques/AML.T0018/"><span class="relation-id">AML.T0018</span><strong>Манипуляция ИИ-моделью</strong><p>Подписание кода дает гарантию, что модель не была изменена после подписания.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><p>Подписание кода дает гарантию, что модель не была изменена после подписания.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><p>Подписание кода дает гарантию, что модель не была изменена после подписания.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.002/"><span class="relation-id">AML.T0018.002</span><strong>Встраивание вредоносного ПО</strong><p>Подписание кода дает гарантию, что модель не была изменена после подписания.</p></a>
</div>
