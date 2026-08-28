---
atlas_id: AML.T0088
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2025-10-31"
description: Злоумышленники могут использовать генеративный искусственный интеллект (GenAI) для создания синтетических медиа, то есть изображений, видео, аудио и текста, которые выглядят подлинными. Такие дипфейки могут...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Enterprise
procedure_count: 4
source_name: Generate Deepfakes
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0001
title: Создание дипфейков
url: /techniques/AML.T0088/
---

Злоумышленники могут использовать генеративный искусственный интеллект (GenAI) для создания синтетических медиа, то есть изображений, видео, аудио и текста, которые выглядят подлинными. Такие [дипфейки]( https://en.wikipedia.org/wiki/Deepfake) могут имитировать реального человека или изображать вымышленных персонажей. Злоумышленники могут использовать дипфейки для имперсонации при проведении [фишинга](/techniques/AML.T0052) или для обхода ИИ-приложений, например систем биометрической проверки личности (см. [Обход ИИ-модели](/techniques/AML.T0015)).

Манипулирование медиа было возможно уже давно, однако GenAI снижает необходимые навыки и трудозатраты, позволяя злоумышленникам быстро масштабировать операции и атаковать больше пользователей или систем. Кроме того, GenAI делает возможными манипуляции в реальном времени.

Злоумышленники могут использовать модели с открытым исходным кодом и ПО, разработанные для легитимных сценариев, чтобы создавать дипфейки для вредоносного использования. Однако существуют и проекты, специально ориентированные на вредоносные сценарии, например [ProKYC](https://www.catonetworks.com/blog/prokyc-selling-deepfake-tool-for-account-fraud-attacks/).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0001/"><span class="relation-id">AML.TA0001</span><strong>Подготовка атаки на ИИ</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0009/"><span class="relation-id">AML.M0009</span><strong>Слияние данных нескольких сенсоров для предиктивного ИИ</strong><p>Использование разных сенсоров, например инфракрасных камер глубины, может помогать обнаруживать дипфейки.</p></a>
<a class="relation-item" href="/mitigations/AML.M0034/"><span class="relation-id">AML.M0034</span><strong>Обнаружение дипфейков</strong><p>Обнаружение дипфейков можно использовать для выявления и блокировки сгенерированного контента.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0033/"><span class="relation-id">AML.CS0033</span><strong>Обход мобильной KYC-верификации с помощью дипфейк-изображения в реальном времени</strong><span class="relation-meta">Актор: iProov Red Team / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Исследователи использовали собранные изображения лица жертвы и инструмент Faceswap, чтобы создать дипфейк-видео в реальном времени, имитирующие внешность жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0034/"><span class="relation-id">AML.CS0034</span><strong>ProKYC: дипфейк-инструмент для атак с мошенническим созданием аккаунтов</strong><span class="relation-meta">Актор: ProKYC, cybercriminal group / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Злоумышленник использовал в ProKYC сочетание реальных персональных данных и поддельной информации, чтобы сгенерировать дипфейк-удостоверение личности.</p></a>
<a class="relation-item" href="/studies/AML.CS0034/"><span class="relation-id">AML.CS0034</span><strong>ProKYC: дипфейк-инструмент для атак с мошенническим созданием аккаунтов</strong><span class="relation-meta">Актор: ProKYC, cybercriminal group / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Злоумышленник использовал ProKYC, чтобы сгенерировать дипфейк-видео с селфи: лицо на видео совпадало с лицом в удостоверении личности, что позволяло обходить проверку присутствия живого человека.</p></a>
<a class="relation-item" href="/studies/AML.CS0057/"><span class="relation-id">AML.CS0057</span><strong>Storm-2139: обход защитных ограничений Azure OpenAI</strong><span class="relation-meta">Актор: Storm-2139 / Тактика: AML.TA0001 Подготовка атаки на ИИ</span><p>Конечные пользователи генерировали синтетические изображения с противоправным содержанием, включая интимные изображения знаменитостей без их согласия, а также другой сексуально откровенный, мизогинный, насильственный или разжигающий ненависть контент.</p></a>
</div>
