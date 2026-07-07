---
atlas_id: AML.TA0003
atlas_type: tactic
attack_ref_id: TA0042
attack_ref_url: https://attack.mitre.org/tactics/TA0042/
created_date: "2022-01-24"
description: Злоумышленник пытается подготовить ресурсы, которые можно использовать для поддержки операций. Подготовка ресурсов включает техники, при которых злоумышленники создают, покупают, компрометируют или крадут ресурсы,...
generated: true
generated_by: atlasgen
modified_date: "2025-04-09"
procedure_count: 101
source_name: Resource Development
technique_count: 39
title: Подготовка ресурсов
url: /tactics/AML.TA0003/
---

Злоумышленник пытается подготовить ресурсы, которые можно использовать для поддержки операций.

Подготовка ресурсов включает техники, при которых злоумышленники создают, покупают, компрометируют или крадут ресурсы, которые могут использоваться для поддержки выбора целей.
Такие ресурсы включают ИИ-артефакты, инфраструктуру, учетные записи или возможности.
Злоумышленник может использовать эти ресурсы на других этапах жизненного цикла атаки, например при [подготовке атаки на ИИ](/tactics/AML.TA0001).


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0002/"><span class="relation-id">AML.T0002</span><strong>Получение публичных ИИ-артефактов</strong></a>
<a class="relation-item" href="/techniques/AML.T0002.000/"><span class="relation-id">AML.T0002.000</span><strong>Наборы данных</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0002.000/"><span class="relation-id">AML.T0002.000</span><strong>Наборы данных</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0002.001/"><span class="relation-id">AML.T0002.001</span><strong>Модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0002.001/"><span class="relation-id">AML.T0002.001</span><strong>Модели</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0002.002/"><span class="relation-id">AML.T0002.002</span><strong>Конфигурация ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0002.002/"><span class="relation-id">AML.T0002.002</span><strong>Конфигурация ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008/"><span class="relation-id">AML.T0008</span><strong>Получение инфраструктуры</strong></a>
<a class="relation-item" href="/techniques/AML.T0008.000/"><span class="relation-id">AML.T0008.000</span><strong>Рабочие пространства для разработки ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.000/"><span class="relation-id">AML.T0008.000</span><strong>Рабочие пространства для разработки ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.001/"><span class="relation-id">AML.T0008.001</span><strong>Потребительское оборудование</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.001/"><span class="relation-id">AML.T0008.001</span><strong>Потребительское оборудование</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.002/"><span class="relation-id">AML.T0008.002</span><strong>Домены</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.002/"><span class="relation-id">AML.T0008.002</span><strong>Домены</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.003/"><span class="relation-id">AML.T0008.003</span><strong>Физические средства противодействия</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.003/"><span class="relation-id">AML.T0008.003</span><strong>Физические средства противодействия</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.004/"><span class="relation-id">AML.T0008.004</span><strong>Serverless-инфраструктура</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.004/"><span class="relation-id">AML.T0008.004</span><strong>Serverless-инфраструктура</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.005/"><span class="relation-id">AML.T0008.005</span><strong>Прокси для ИИ-сервисов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0008.005/"><span class="relation-id">AML.T0008.005</span><strong>Прокси для ИИ-сервисов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016/"><span class="relation-id">AML.T0016</span><strong>Получение средств для атаки</strong></a>
<a class="relation-item" href="/techniques/AML.T0016.000/"><span class="relation-id">AML.T0016.000</span><strong>Готовые реализации состязательных атак на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.000/"><span class="relation-id">AML.T0016.000</span><strong>Готовые реализации состязательных атак на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.001/"><span class="relation-id">AML.T0016.001</span><strong>Программные инструменты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.001/"><span class="relation-id">AML.T0016.001</span><strong>Программные инструменты</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.002/"><span class="relation-id">AML.T0016.002</span><strong>Генеративный ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0016.002/"><span class="relation-id">AML.T0016.002</span><strong>Генеративный ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0017/"><span class="relation-id">AML.T0017</span><strong>Разработка средств для атаки</strong></a>
<a class="relation-item" href="/techniques/AML.T0017.000/"><span class="relation-id">AML.T0017.000</span><strong>Состязательные атаки на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0017.000/"><span class="relation-id">AML.T0017.000</span><strong>Состязательные атаки на ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0019/"><span class="relation-id">AML.T0019</span><strong>Публикация отравленных наборов данных</strong></a>
<a class="relation-item" href="/techniques/AML.T0020/"><span class="relation-id">AML.T0020</span><strong>Отравление обучающих данных</strong></a>
<a class="relation-item" href="/techniques/AML.T0021/"><span class="relation-id">AML.T0021</span><strong>Создание учетных записей</strong></a>
<a class="relation-item" href="/techniques/AML.T0058/"><span class="relation-id">AML.T0058</span><strong>Публикация отравленных моделей</strong></a>
<a class="relation-item" href="/techniques/AML.T0060/"><span class="relation-id">AML.T0060</span><strong>Публикация галлюцинированных сущностей</strong></a>
<a class="relation-item" href="/techniques/AML.T0065/"><span class="relation-id">AML.T0065</span><strong>Создание промптов для LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0066/"><span class="relation-id">AML.T0066</span><strong>Подготовка содержимого для извлечения</strong></a>
<a class="relation-item" href="/techniques/AML.T0079/"><span class="relation-id">AML.T0079</span><strong>Размещение средств атаки</strong></a>
<a class="relation-item" href="/techniques/AML.T0104/"><span class="relation-id">AML.T0104</span><strong>Публикация отравленного инструмента ИИ-агента</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0000/"><span class="relation-id">AML.CS0000</span><strong>Обход детектора C&amp;C-трафика вредоносного ПО на основе глубокого обучения</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Мы получили набор данных HTTP-трафика командования и управления, состоящий примерно из 33 млн безвредных и 27 млн вредоносных заголовков HTTP-пакетов.</p></a>
<a class="relation-item" href="/studies/AML.CS0001/"><span class="relation-id">AML.CS0001</span><strong>Обход обнаружения DGA-доменов ботнетов</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи получили общедоступную CNN-модель обнаружения DGA и протестировали ее на известном наборе данных доменных имен, сгенерированных DGA, который включает около 50 млн доменных имен из 64 семейств ботнетных DGA. CNN-модель обнаружения DGA показала точность обнаружения выше 70% на 16, то есть примерно 25%, семействах ботнетных DGA.</p></a>
<a class="relation-item" href="/studies/AML.CS0001/"><span class="relation-id">AML.CS0001</span><strong>Обход обнаружения DGA-доменов ботнетов</strong><span class="relation-meta">Актор: Palo Alto Networks AI Research Team / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи разработали универсальную технику мутации, требующую минимального числа итераций.</p></a>
<a class="relation-item" href="/studies/AML.CS0002/"><span class="relation-id">AML.CS0002</span><strong>Отравление VirusTotal</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленник получил [metame](https://github.com/a0rtega/metame), простой движок метаморфного кода для произвольных исполняемых файлов.</p></a>
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи использовали информацию о репутационном скоринге, чтобы с помощью обратной разработки определить, какие атрибуты давали тот или иной уровень положительной или отрицательной репутации. В процессе они обнаружили вторичную модель, которая переопределяла первую модель. Положительные оценки вторичной модели переопределяли решение основной ML-модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленники приобрели кастомизированные недорогие мобильные телефоны.</p></a>
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленники получили ПО, которое преобразует статичные фотографии в видео, добавляя реалистичные эффекты, например моргание.</p></a>
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленники получили кастомизированные Android ROM и приложение виртуальной камеры.</p></a>
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленники использовали идентификационные данные жертв для регистрации новых учетных записей в налоговой системе.</p></a>
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи собрали похожие наборы данных, которые использовали целевые сервисы машинного перевода.</p></a>
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи собрали похожие архитектуры моделей, которые использовали целевые сервисы машинного перевода.</p></a>
<a class="relation-item" href="/studies/AML.CS0006/"><span class="relation-id">AML.CS0006</span><strong>Ошибочная конфигурация Clearview AI</strong><span class="relation-meta">Актор: Researchers at spiderSilk / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Злоумышленники могли скачать обучающие данные и извлечь из исходного кода и декомпилированных бинарных файлов приложений сведения о ПО, моделях и возможностях системы.</p></a>
</div>


Показано 12 из 101 примеров.
