---
atlas_id: AML.T0087
atlas_type: technique
attack_ref_id: T1589
attack_ref_url: https://attack.mitre.org/techniques/T1589/
created_date: "2025-10-31"
description: 'Злоумышленники могут собирать сведения о личности жертвы, которые затем используются при выборе цели и подготовке атаки. Такие сведения могут включать разные данные: персональные данные (например, имена сотрудников,...'
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 3
source_name: Gather Victim Identity Information
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0002
title: Сбор сведений о личности жертвы
url: /techniques/AML.T0087/
---

Злоумышленники могут собирать сведения о личности жертвы, которые затем используются при выборе цели и подготовке атаки. Такие сведения могут включать разные данные: персональные данные (например, имена сотрудников, адреса электронной почты, фотографии и т. д.), а также чувствительную информацию, например учетные данные или настройки многофакторной аутентификации (MFA).

Злоумышленники могут собирать эту информацию разными способами: через прямое выманивание, [поиск на сайтах, принадлежащих жертве](/techniques/AML.T0003), или через утечки информации на черном рынке.

Собранные данные жертвы могут использоваться для создания дипфейков и убедительной имитации личности жертвы. Это может помочь злоумышленникам [создавать учетные записи](/techniques/AML.T0021) от имени имперсонируемого лица или проводить убедительные [фишинговые](/techniques/AML.T0052) атаки.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0002 Разведка</span><p>Злоумышленники собрали идентификационные данные пользователей и фотографии лиц в высоком разрешении на онлайн-черном рынке.</p></a>
<a class="relation-item" href="/studies/AML.CS0033/"><span class="relation-id">AML.CS0033</span><strong>Обход мобильной KYC-верификации с помощью дипфейк-изображения в реальном времени</strong><span class="relation-meta">Актор: iProov Red Team / Тактика: AML.TA0002 Разведка</span><p>Исследователи собрали идентификационные данные пользователей и изображения лиц в высоком разрешении из онлайн-соцсетей и/или с площадок черного рынка.</p></a>
<a class="relation-item" href="/studies/AML.CS0034/"><span class="relation-id">AML.CS0034</span><strong>ProKYC: дипфейк-инструмент для атак с мошенническим созданием аккаунтов</strong><span class="relation-meta">Актор: ProKYC, cybercriminal group / Тактика: AML.TA0002 Разведка</span><p>Злоумышленник собрал идентификационные данные пользователя.</p></a>
</div>
