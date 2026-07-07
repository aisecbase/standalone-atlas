---
atlas_id: AML.T0002.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут получать публичные модели для использования в своих операциях. Злоумышленники могут искать модели, используемые организацией-жертвой, или модели, репрезентативные по отношению к моделям, которые...
generated: true
generated_by: atlasgen
maturity: demonstrated
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Predictive AI
    - Generative AI
    - Agentic AI
procedure_count: 4
source_name: Models
subtechnique_count: 0
subtechnique_of: AML.T0002
tactics:
    - AML.TA0003
title: Модели
url: /techniques/AML.T0002.001/
---

Злоумышленники могут получать публичные модели для использования в своих операциях.
Злоумышленники могут искать модели, используемые организацией-жертвой, или модели, репрезентативные по отношению к моделям, которые использует организация-жертва.
Репрезентативные модели могут включать архитектуры моделей или предварительно обученные модели, которые задают архитектуру, а также параметры модели, полученные при обучении на наборе данных.
Злоумышленник может искать в публичных источниках распространенные форматы файлов конфигурации архитектуры модели, такие как YAML или конфигурационные файлы Python, а также распространенные форматы файлов хранения моделей, такие как ONNX (`.onnx`), HDF5 (`.h5`), Pickle (`.pkl`), PyTorch (`.pth`) или TensorFlow (`.pb`, `.tflite`).

Полученные модели полезны для развития операций злоумышленника и часто используются для адаптации атак под модель жертвы.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0003/"><span class="relation-id">AML.TA0003</span><strong>Подготовка ресурсов</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0002/"><span class="relation-id">AML.T0002</span><strong>Получение публичных ИИ-артефактов</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0002.000/"><span class="relation-id">AML.T0002.000</span><strong>Наборы данных</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0002.002/"><span class="relation-id">AML.T0002.002</span><strong>Конфигурация ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0001/"><span class="relation-id">AML.M0001</span><strong>Ограничение публикации артефактов модели</strong><p>Ограничение публикации архитектур моделей и контрольных точек может снизить способность злоумышленника нацеливаться на эти модели.</p></a>
<a class="relation-item" href="/mitigations/AML.M0014/"><span class="relation-id">AML.M0014</span><strong>Проверка ИИ-артефактов</strong><p>Внедрите надлежащую проверку подписей, чтобы небезопасные ИИ-модели не попадали в систему.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0005/"><span class="relation-id">AML.CS0005</span><strong>Атака на сервисы машинного перевода</strong><span class="relation-meta">Актор: Berkeley Artificial Intelligence Research / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи собрали похожие архитектуры моделей, которые использовали целевые сервисы машинного перевода.</p></a>
<a class="relation-item" href="/studies/AML.CS0007/"><span class="relation-id">AML.CS0007</span><strong>Репликация модели GPT-2</strong><span class="relation-meta">Актор: Researchers at Brown University / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи получили эталонную реализацию похожей общедоступной модели под названием Grover.</p></a>
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи получили APK-файлы приложений из Google Play. Они отфильтровали список потенциальных целевых приложений, проверяя метаданные кода на ключевые слова, связанные с TensorFlow или TFLite, а также с их бинарными форматами моделей (.tf и .tflite). Модели были извлечены из APK-файлов с помощью Apktool.</p></a>
<a class="relation-item" href="/studies/AML.CS0019/"><span class="relation-id">AML.CS0019</span><strong>PoisonGPT</strong><span class="relation-meta">Актор: Mithril Security Researchers / Тактика: AML.TA0003 Подготовка ресурсов</span><p>Исследователи загрузили модель [GPT-J-6B с Hugging Face](https://huggingface.co/EleutherAI/gpt-j-6b) с открытым исходным кодом. GPT-J-6B — это большая языковая модель, обычно используемая для генерации текста по входным промптам, например в задачах ответов на вопросы.</p></a>
</div>
