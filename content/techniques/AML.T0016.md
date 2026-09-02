---
atlas_id: AML.T0016
atlas_type: technique
attack_ref_id: T1588
attack_ref_url: https://attack.mitre.org/techniques/T1588/
created_date: "2021-05-13"
description: Злоумышленники могут искать и получать программные средства для использования в своих операциях. Такие средства могут быть специфичны для атак на основе ИИ, например готовые реализации состязательных атак на ИИ, либо...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
    - Enterprise
procedure_count: 1
source_name: Obtain Capabilities
subtechnique_count: 5
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Получение средств для атаки
url: /techniques/AML.T0016/
---

Злоумышленники могут искать и получать программные средства для использования в своих операциях.

Такие средства могут быть специфичны для атак на основе ИИ, например [готовые реализации состязательных атак на ИИ](/techniques/AML.T0016.000), либо представлять собой универсальные программные инструменты, переиспользуемые во вредоносных целях ([программные инструменты](/techniques/AML.T0016.001)). В обоих случаях злоумышленник может изменить или адаптировать такое средство, чтобы упростить нацеливание на конкретную систему с поддержкой ИИ.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0016.000/"><span class="relation-id">AML.T0016.000</span><strong>Готовые реализации состязательных атак на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.001/"><span class="relation-id">AML.T0016.001</span><strong>Программные инструменты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.002/"><span class="relation-id">AML.T0016.002</span><strong>Генеративный ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.003/"><span class="relation-id">AML.T0016.003</span><strong>Exploits</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.004/"><span class="relation-id">AML.T0016.004</span><strong>AI Agent Tools</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0001/"><span class="relation-id">AML.M0001</span><strong>Ограничение публикации артефактов модели</strong><p>Ограничивайте публикацию в открытом доступе артефактов моделей, используемых в продакшене: злоумышленники могут получить такие артефакты и адаптировать их для использования в своих операциях.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0033/"><span class="relation-id">AML.CS0033</span><strong>Обход мобильной KYC-верификации с помощью дипфейк-изображения в реальном времени</strong><span class="relation-meta">Актор: iProov Red Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи получили [Virtual Camera: Live Assist](https://apkpure.com/virtual-camera-live-assist/virtual.camera.app) — Android-приложение, позволяющее заменить камеру устройства видеопотоком. Приложение работает на настоящих Android-устройствах без root-доступа.</p></a>
</div>
