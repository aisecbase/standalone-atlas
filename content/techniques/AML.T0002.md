---
atlas_id: AML.T0002
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут искать в публичных источниках, включая облачные хранилища, публично доступные сервисы, а также репозитории ПО или данных, чтобы выявить ИИ-артефакты. Такие ИИ-артефакты могут включать программный...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 3
source_name: Acquire Public AI Artifacts
subtechnique_count: 3
subtechnique_of: ""
tactics:
    - AML.TA0003
title: Получение публичных ИИ-артефактов
url: /techniques/AML.T0002/
---

Злоумышленники могут искать в публичных источниках, включая облачные хранилища, публично доступные сервисы, а также репозитории ПО или данных, чтобы выявить ИИ-артефакты.
Такие ИИ-артефакты могут включать программный стек, используемый для обучения и развертывания моделей, обучающие и тестовые данные, конфигурации и параметры моделей.
Злоумышленника будут особенно интересовать артефакты, размещенные организацией-жертвой или связанные с ней, поскольку они могут отражать то, что организация использует в продакшене.
Злоумышленники могут выявлять репозитории артефактов через другие ресурсы, связанные с организацией-жертвой, например [поиск на сайтах организации-жертвы](/techniques/AML.T0003) или [поиск в открытых технических базах данных](/techniques/AML.T0000).
Такие ИИ-артефакты часто дают злоумышленникам сведения о задаче ИИ и используемом подходе.

ИИ-артефакты могут помочь злоумышленнику [создать прокси-модель ИИ](/techniques/AML.T0005).
Если эти артефакты включают части реальной модели, используемой в продакшене, их можно напрямую использовать для [создания состязательных данных](/techniques/AML.T0043).
Получение некоторых артефактов требует регистрации с указанием пользовательских данных, например email или имени, ключей AWS или письменных запросов, а также может потребовать от злоумышленника [создать учетные записи](/techniques/AML.T0021).

Артефакты могут размещаться на инфраструктуре, контролируемой жертвой, что дает жертве некоторую информацию о том, кто обращался к этим данным.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Подтехники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0002.000/"><span class="relation-id">AML.T0002.000</span><strong>Наборы данных</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0002.001/"><span class="relation-id">AML.T0002.001</span><strong>Модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0002.002/"><span class="relation-id">AML.T0002.002</span><strong>Конфигурация ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Ограничьте публикацию чувствительной информации в метаданных развернутых систем и публично доступных приложений.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0001/"><span class="relation-id">AML.CS0001</span><strong>Обход обнаружения DGA-доменов ботнетов</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи получили общедоступную CNN-модель обнаружения DGA и протестировали ее на известном наборе данных доменных имен, сгенерированных DGA, который включает около 50 млн доменных имен из 64 семейств ботнетных DGA. CNN-модель обнаружения DGA показала точность обнаружения выше 70% на 16, то есть примерно 25%, семействах ботнетных DGA.</p></a>
<a class="relation-item" href="/studies/AML.CS0006/"><span class="relation-id">AML.CS0006</span><strong>Ошибочная конфигурация Clearview AI</strong><span class="relation-meta">Актор: Researchers at spiderSilk / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленники могли скачать обучающие данные и извлечь из исходного кода и декомпилированных бинарных файлов приложений сведения о ПО, моделях и возможностях системы.</p></a>
<a class="relation-item" href="/studies/AML.CS0011/"><span class="relation-id">AML.CS0011</span><strong>Обход ИИ на периферии Microsoft</strong><span class="relation-meta">Актор: Azure Red Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Команда выявила и получила общедоступную базовую модель, чтобы использовать ее против целевой ML-модели.</p></a>
</div>
