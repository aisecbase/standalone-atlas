---
atlas_id: AML.T0021
atlas_type: technique
attack_ref_id: T1585
attack_ref_url: https://attack.mitre.org/techniques/T1585/
created_date: "2022-01-24"
description: Злоумышленники могут создавать учетные записи в различных сервисах, чтобы использовать их при выборе целей, получить доступ к ресурсам, необходимым для подготовки атаки на ИИ, или выдавать себя за жертву.
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 6
source_name: Establish Accounts
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Создание учетных записей
url: /techniques/AML.T0021/
---

Злоумышленники могут создавать учетные записи в различных сервисах, чтобы использовать их при выборе целей, получить доступ к ресурсам, необходимым для [подготовки атаки на ИИ](/tactics/AML.TA0001), или выдавать себя за жертву.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Проверяйте идентичность субъекта перед предоставлением доступа к реестрам моделей, чтобы вновь созданные учётные записи не получали автоматического доступа к защищённым ИИ-артефактам.</p></a>
<a class="relation-item" href="/mitigations/AML.M0019/"><span class="relation-id">AML.M0019</span><strong>Контроль доступа к ИИ-моделям и данным в продакшене</strong><p>Verify identities before granting production AI access so newly established accounts cannot automatically access protected AI services.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленники использовали идентификационные данные жертв для регистрации новых учетных записей в налоговой системе.</p></a>
<a class="relation-item" href="/studies/AML.CS0006/"><span class="relation-id">AML.CS0006</span><strong>Ошибочная конфигурация Clearview AI</strong><span class="relation-meta">Актор: Researchers at spiderSilk / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь безопасности получил первоначальный доступ к приватному репозиторию кода Clearview AI из-за ошибочной настройки сервера, которая позволяла произвольному пользователю зарегистрировать действительную учетную запись.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователь зарегистрировал на Hugging Face непроверенную учетную запись «организации», заняв пространство имен целевой компании.</p></a>
<a class="relation-item" href="/studies/AML.CS0033/"><span class="relation-id">AML.CS0033</span><strong>Обход мобильной KYC-верификации с помощью дипфейк-изображения в реальном времени</strong><span class="relation-meta">Актор: iProov Red Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи использовали собранные сведения о жертве, чтобы зарегистрировать аккаунт в приложении финансового сервиса.</p></a>
<a class="relation-item" href="/studies/AML.CS0034/"><span class="relation-id">AML.CS0034</span><strong>ProKYC: дипфейк-инструмент для атак с мошенническим созданием аккаунтов</strong><span class="relation-meta">Актор: ProKYC, cybercriminal group / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник использовал сведения о жертве, чтобы зарегистрировать аккаунт в приложении финансового сервиса, например на криптовалютной бирже.</p></a>
<a class="relation-item" href="/studies/AML.CS0065/"><span class="relation-id">AML.CS0065</span><strong>Атака на цепочку поставок через повторное использование пространства имён модели</strong><span class="relation-meta">Актор: Unit 42 Researchers / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Специалисты Unit 42 зарегистрировали организацию в Hugging Face, использовав заброшенное пространство имён, связанное с путём к модели, которой ранее доверяли. Учётная запись первоначального владельца не была скомпрометирована.</p></a>
</div>
