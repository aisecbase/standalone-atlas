---
atlas_id: AML.T0004
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут искать в открытых репозиториях приложений при выборе целей. Примеры таких репозиториев включают Google Play, iOS App Store, macOS App Store и Microsoft Store. Злоумышленники могут составлять...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 3
source_name: Search Application Repositories
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0002
title: Поиск в репозиториях приложений
url: /techniques/AML.T0004/
---

Злоумышленники могут искать в открытых репозиториях приложений при выборе целей.
Примеры таких репозиториев включают Google Play, iOS App Store, macOS App Store и Microsoft Store.

Злоумышленники могут составлять поисковые запросы для поиска приложений, содержащих компоненты с поддержкой ИИ.
Часто следующим шагом становится [получение публичных ИИ-артефактов](/techniques/AML.T0002).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0002/"><span class="relation-id">AML.TA0002</span><strong>Разведка</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0000/"><span class="relation-id">AML.M0000</span><strong>Ограничение публичного раскрытия информации</strong><p>Ограничьте публикацию чувствительной информации в метаданных развернутых систем и публично доступных приложений.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0002 Разведка</span><p>Чтобы составить список потенциальных целевых моделей, исследователи искали в Google Play приложения, которые могли содержать встроенные модели глубокого обучения, по ключевым словам, связанным с глубоким обучением.</p></a>
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0002 Разведка</span><p>Исследователи Trend Micro использовали сервисы индексирования и инструменты веб-поиска, чтобы выявить более 8 000 приватных реестров контейнеров, доступных из интернета. Примерно 70% реестров имели чрезмерно широкие права доступа, разрешавшие запись. Среди приватных реестров контейнеров были как самостоятельно размещенные реестры, так и реестры, развернутые у поставщиков облачных услуг (CSP). Реестры оказались доступны из-за сочетания следующих факторов:</p><ul><li>ошибка конфигурации, открывшая публичный доступ к приватному реестру;</li><li>отсутствие корректных механизмов аутентификации и авторизации;</li><li>недостаточная сегментация сети и слабые средства контроля доступа.</li></ul></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0002 Разведка</span><p>Исследователи искали цели среди приложений, которые, вероятно, построены на LLM-фреймворках и могут использовать функции, уязвимые к RCE. Для этого они сканировали репозитории исходного кода в поисках URL развернутых приложений.</p></a>
</div>
