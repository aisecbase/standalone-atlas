---
atlas_id: AML.T0010.003
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Системы с поддержкой ИИ часто по-разному опираются на модели с открытым исходным кодом. Чаще всего организация-жертва может использовать такие модели для дообучения. Эти модели загружаются из внешнего источника и...
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 6
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 6
source_name: Model
subtechnique_count: 0
subtechnique_of: AML.T0010
tactics:
    - AML.TA0004
title: Модель
url: /techniques/AML.T0010.003/
---

Системы с поддержкой ИИ часто по-разному опираются на модели с открытым исходным кодом.

Чаще всего организация-жертва может использовать такие модели для дообучения.

Эти модели загружаются из внешнего источника и затем используются как основа модели при ее дообучении на меньшем закрытом наборе данных.

Загрузка моделей часто требует выполнения сохраненного кода в виде сохраненного файла модели.

Такие модели могут быть скомпрометированы с помощью традиционного вредоносного ПО или некоторых техник состязательных атак на ИИ.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010/"><span class="relation-id">AML.T0010</span><strong>Компрометация цепочки поставок ИИ</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010.000/"><span class="relation-id">AML.T0010.000</span><strong>Аппаратное обеспечение</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.001/"><span class="relation-id">AML.T0010.001</span><strong>ПО для ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.002/"><span class="relation-id">AML.T0010.002</span><strong>Данные</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.004/"><span class="relation-id">AML.T0010.004</span><strong>Реестр контейнеров</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.005/"><span class="relation-id">AML.T0010.005</span><strong>Инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа может предотвращать подмену ML-артефактов и несанкционированное копирование.</p></a>
<a class="relation-item" href="/mitigations/AML.M0006/"><span class="relation-id">AML.M0006</span><strong>Использование ансамблевых методов</strong><p>Использование нескольких разных моделей обеспечивает минимальную потерю производительности, если уязвимость обнаружена в инструменте для одной модели или семейства моделей.</p></a>
<a class="relation-item" href="/mitigations/AML.M0008/"><span class="relation-id">AML.M0008</span><strong>Валидация ИИ-модели</strong><p>Убедитесь, что приобретенные модели не реагируют на потенциальные бэкдор-триггеры или состязательное воздействие.</p></a>
<a class="relation-item" href="/mitigations/AML.M0013/"><span class="relation-id">AML.M0013</span><strong>Подписание кода</strong><p>Требуйте корректной подписи файлов модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0017/"><span class="relation-id">AML.M0017</span><strong>Методы распространения ИИ-моделей</strong><p>Злоумышленник может переупаковать приложение с вредоносной версией модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0035/"><span class="relation-id">AML.M0035</span><strong>Красная команда по ИИ</strong><p>Внедряйте недоверенную или модифицированную модель только в контролируемую среду через репрезентативные пути её получения и развёртывания. Проверьте сведения о происхождении, сканирование, подписание, процедуры одобрения, изоляцию и возможность возврата к предыдущей версии модели.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0004 Первичный доступ</span><p>На практике вредоносный APK-файл должен быть установлен на устройства жертв через компрометацию цепочки поставок.</p></a>
<a class="relation-item" href="/studies/AML.CS0019/"><span class="relation-id">AML.CS0019</span><strong>PoisonGPT</strong><span class="relation-meta">Актор: Mithril Security Researchers / Тактика: AML.TA0004 Первичный доступ</span><p>Ничего не подозревающие пользователи могли бы скачать состязательную модель и интегрировать её в приложения. После того как исследователи сообщили о проведённом испытании, Hugging Face отключила репозиторий с похожим именем.</p></a>
<a class="relation-item" href="/studies/AML.CS0023/"><span class="relation-id">AML.CS0023</span><strong>ShadowRay</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0004 Первичный доступ</span><p>Токены Hugging Face могли позволить злоумышленнику заменить модели организации-жертвы вредоносными вариантами.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0004 Первичный доступ</span><p>Цепочка поставок ИИ-моделей жертвы оказалась скомпрометирована. Пользователи репозитория моделей будут получать модель злоумышленника со встроенным вредоносным ПО.</p></a>
<a class="relation-item" href="/studies/AML.CS0064/"><span class="relation-id">AML.CS0064</span><strong>Отравленные шаблоны GGUF: атака на цепочку поставок во время инференса</strong><span class="relation-meta">Актор: Pillar Security, Fujitsu Research of Europe / Тактика: AML.TA0004 Первичный доступ</span><p>Жертва скачивает и интегрирует отравленный артефакт, доверяя модели и включённым в её состав компонентам. В результате бэкдор в шаблоне внедряется в ИИ-приложение или ИИ-агента жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0065/"><span class="relation-id">AML.CS0065</span><strong>Атака на цепочку поставок через повторное использование пространства имён модели</strong><span class="relation-meta">Актор: Unit 42 Researchers / Тактика: AML.TA0004 Первичный доступ</span><p>В облачном каталоге моделей, приложении или пайплайне развёртывания разрешение сохранявшейся без обновления ссылки на модель, заданной только именем без фиксации ревизии, приводило к выбору подконтрольной злоумышленнику модели-замены. Специалисты Unit 42 продемонстрировали это на примере развёртываний в Vertex AI и Azure AI Foundry.</p></a>
</div>
