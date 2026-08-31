---
atlas_id: AML.T0052
atlas_type: technique
attack_ref_id: T1566
attack_ref_url: https://attack.mitre.org/techniques/T1566/
created_date: "2023-10-25"
description: Злоумышленники могут отправлять фишинговые сообщения, чтобы получить доступ к системам организации-жертвы. Все формы фишинга представляют собой социальную инженерию, доставляемую по электронным каналам. Фишинг может...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 2
source_name: Phishing
subtechnique_count: 2
subtechnique_of: ""
tactics:
    - AML.TA0004
    - AML.TA0015
title: Фишинг
url: /techniques/AML.T0052/
---

Злоумышленники могут отправлять фишинговые сообщения, чтобы получить доступ к системам организации-жертвы. Все формы фишинга представляют собой социальную инженерию, доставляемую по электронным каналам. Фишинг может быть целевым; такой подход известен как целевой фишинг. При целевом фишинге злоумышленник нацелен на конкретного человека, компанию или отрасль. В более общем случае злоумышленники могут проводить нецелевой фишинг, например в массовых спам-кампаниях с вредоносным ПО.

Генеративный ИИ, включая LLM, которые создают синтетический текст, визуальные дипфейки лиц и аудиодипфейки речи (см. [Создание дипфейков](/techniques/AML.T0088)), позволяет злоумышленникам масштабировать целевые фишинговые кампании (см. [целевой фишинг через LLM для социальной инженерии](/techniques/AML.T0052.000)). LLM могут взаимодействовать с пользователями в текстовых диалогах и могут быть настроены системным промптом на выманивание чувствительной информации. Дипфейки также могут использоваться при [имперсонации](/techniques/AML.T0073) как вспомогательное средство для фишинга.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
<a class="relation-item" href="/tactics/AML.TA0015/"><span class="relation-id">AML.TA0015</span><strong>Латеральное перемещение</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0052.000/"><span class="relation-id">AML.T0052.000</span><strong>Целевой фишинг через LLM для социальной инженерии</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0052.001/"><span class="relation-id">AML.T0052.001</span><strong>Фишинг с использованием дипфейков</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0018/"><span class="relation-id">AML.M0018</span><strong>Обучение пользователей</strong><p>Обучайте пользователей распознавать фишинговые попытки злоумышленника, чтобы снизить риск успешного целевого фишинга, социальной инженерии и других техник с участием пользователя.</p></a>
<a class="relation-item" href="/mitigations/AML.M0034/"><span class="relation-id">AML.M0034</span><strong>Обнаружение дипфейков</strong><p>Обнаружение дипфейков можно использовать для выявления и блокировки фишинговых попыток, использующих сгенерированный контент.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0032/"><span class="relation-id">AML.CS0032</span><strong>Попытка обхода ML-системы обнаружения фишинговых веб-страниц</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0004 Первичный доступ</span><p>Если злоумышленнику удается успешно избежать обнаружения, он может продолжать эксплуатацию фишинговых сайтов и красть учетные данные жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее команды, динамически генерируемые ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0015 Латеральное перемещение</span><p>APT28 отправила с этой учетной записи фишинговое письмо с вложением, содержащим вредоносное ПО.</p></a>
</div>
