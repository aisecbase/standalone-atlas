---
atlas_id: AML.TA0008
atlas_type: tactic
attack_ref_id: TA0007
attack_ref_url: https://attack.mitre.org/tactics/TA0007/
created_date: "2022-01-24"
description: Злоумышленник пытается разобраться в ИИ-среде организации-жертвы. Выявление включает техники, которые злоумышленник может использовать для получения сведений о системе и внутренней сети. Эти техники помогают...
generated: true
generated_by: atlasgen
modified_date: "2025-04-09"
procedure_count: 19
source_name: Discovery
technique_count: 23
title: Выявление
url: /tactics/AML.TA0008/
---

Злоумышленник пытается разобраться в ИИ-среде организации-жертвы.

Выявление включает техники, которые злоумышленник может использовать для получения сведений о системе и внутренней сети. Эти техники помогают злоумышленникам наблюдать за средой и ориентироваться в ней, прежде чем решить, как действовать дальше. Они также позволяют злоумышленникам исследовать, что они могут контролировать и что находится вокруг их точки входа, чтобы понять, как это может помочь достижению текущей цели. Для такого сбора информации после компрометации часто используются встроенные средства операционной системы.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0007/"><span class="relation-id">AML.T0007</span><strong>Выявление ИИ-артефактов</strong></a>
<a class="relation-item" href="/techniques/AML.T0013/"><span class="relation-id">AML.T0013</span><strong>Выявление онтологии ИИ-модели</strong></a>
<a class="relation-item" href="/techniques/AML.T0014/"><span class="relation-id">AML.T0014</span><strong>Выявление семейства ИИ-модели</strong></a>
<a class="relation-item" href="/techniques/AML.T0062/"><span class="relation-id">AML.T0062</span><strong>Выявление галлюцинированных сущностей LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0063/"><span class="relation-id">AML.T0063</span><strong>Выявление выходных данных ИИ-модели</strong></a>
<a class="relation-item" href="/techniques/AML.T0069/"><span class="relation-id">AML.T0069</span><strong>Выявление системной информации LLM</strong></a>
<a class="relation-item" href="/techniques/AML.T0069.000/"><span class="relation-id">AML.T0069.000</span><strong>Наборы специальных символов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0069.000/"><span class="relation-id">AML.T0069.000</span><strong>Наборы специальных символов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0069.001/"><span class="relation-id">AML.T0069.001</span><strong>Ключевые слова системных инструкций</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0069.001/"><span class="relation-id">AML.T0069.001</span><strong>Ключевые слова системных инструкций</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0069.002/"><span class="relation-id">AML.T0069.002</span><strong>Системный промпт</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0069.002/"><span class="relation-id">AML.T0069.002</span><strong>Системный промпт</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0075/"><span class="relation-id">AML.T0075</span><strong>Выявление облачных сервисов</strong></a>
<a class="relation-item" href="/techniques/AML.T0084/"><span class="relation-id">AML.T0084</span><strong>Выявление конфигурации ИИ-агента</strong></a>
<a class="relation-item" href="/techniques/AML.T0084.000/"><span class="relation-id">AML.T0084.000</span><strong>Встроенные знания</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.000/"><span class="relation-id">AML.T0084.000</span><strong>Встроенные знания</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.001/"><span class="relation-id">AML.T0084.001</span><strong>Определения инструментов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.001/"><span class="relation-id">AML.T0084.001</span><strong>Определения инструментов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.002/"><span class="relation-id">AML.T0084.002</span><strong>Триггеры активации</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.002/"><span class="relation-id">AML.T0084.002</span><strong>Триггеры активации</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.003/"><span class="relation-id">AML.T0084.003</span><strong>Цепочки вызовов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0084.003/"><span class="relation-id">AML.T0084.003</span><strong>Цепочки вызовов</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0089/"><span class="relation-id">AML.T0089</span><strong>Выявление процессов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0003/"><span class="relation-id">AML.CS0003</span><strong>Обход ИИ-детектора вредоносного ПО Cylance</strong><span class="relation-meta">Актор: Skylight Cyber / Тактика: AML.TA0008 Выявление</span><p>Исследователи включили подробное логирование, раскрывающее внутреннюю логику работы ML-модели, особенно в части репутационного скоринга и ансамблирования моделей.</p></a>
<a class="relation-item" href="/studies/AML.CS0008/"><span class="relation-id">AML.CS0008</span><strong>Обход ProofPoint</strong><span class="relation-meta">Актор: Researchers at Silent Break Security / Тактика: AML.TA0008 Выявление</span><p>Исследователи обнаружили, что ProofPoint Email Protection оставляла выходные оценки модели в заголовках писем.</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0008 Выявление</span><p>Команда определила список идентичностей, на которые была нацелена модель, отправляя запросы к API инференса целевой модели.</p></a>
<a class="relation-item" href="/studies/AML.CS0022/"><span class="relation-id">AML.CS0022</span><strong>Галлюцинация пакетов ChatGPT</strong><span class="relation-meta">Актор: Vulcan Cyber, Lasso Security / Тактика: AML.TA0008 Выявление</span><p>Исследователи просили ChatGPT предложить программные пакеты и выявляли среди рекомендаций галлюцинации — пакеты, которых нет в публичном репозитории. Например, на вопрос &#34;how to upload a model to huggingface?&#34; модель предложила установить пакет `huggingface-cli` командой `pip install huggingface-cli`. Такого пакета в PyPI не существовало; реальный CLI-инструмент Hugging Face входит в пакет `huggingface_hub`.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0008 Выявление</span><p>Исследуя ответы Copilot, исследователи выявили специальные разделители и маркеры, например `**`, `**END**`, `Actual Snippet:` и `[^1^]`. Эти строки используются как служебные признаки для отделения разных частей промпта Copilot друг от друга.</p></a>
<a class="relation-item" href="/studies/AML.CS0026/"><span class="relation-id">AML.CS0026</span><strong>Перехват финансовой транзакции с использованием M365 Copilot в роли инсайдера</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0008 Выявление</span><p>Исследуя ответы Copilot, исследователи выявили плагины и конкретные функции, к которым Copilot имеет доступ. Среди них были функция `search_enterprise` и объект `EmailMessage`.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0008 Выявление</span><p>Исследователь мог искать ИИ-модели в среде организации-жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0028/"><span class="relation-id">AML.CS0028</span><strong>Подмена ИИ-модели через атаку на цепочку поставок</strong><span class="relation-meta">Актор: Trend Micro Nebula Cloud Research Team / Тактика: AML.TA0008 Выявление</span><p>Исследователи обнаружили 1 453 уникальные ИИ-модели, встроенные в приватные контейнерные образы. Около половины из них были в формате Open Neural Network Exchange (ONNX).</p></a>
<a class="relation-item" href="/studies/AML.CS0030/"><span class="relation-id">AML.CS0030</span><strong>LLM-джекинг</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0008 Выявление</span><p>Злоумышленники использовали keychecker, чтобы выяснить, какие LLM-сервисы включены в облачной среде и есть ли для этих сервисов квоты ресурсов. Затем злоумышленники проверили, дают ли украденные учетные данные доступ к LLM-ресурсам. Они использовали легитимные запросы `invokeModel` с недопустимым значением -1 для параметра `max_tokens_to_sample`: если учетные данные не давали нужного доступа для вызова модели, такой запрос вызывал ошибку `AccessDenied`. Проверка показала, что украденные учетные данные действительно предоставляли доступ к LLM-ресурсам. Злоумышленники также использовали `GetModelInvocationLoggingConfiguration`, чтобы понять, как настроена модель. Это позволяло им определить, включено ли логирование промптов, и избегать обнаружения при выполнении промптов.</p></a>
<a class="relation-item" href="/studies/AML.CS0036/"><span class="relation-id">AML.CS0036</span><strong>AIKatz: атака на десктопные LLM-приложения</strong><span class="relation-meta">Актор: Lumia Security / Тактика: AML.TA0008 Выявление</span><p>Злоумышленник получил список всех процессов, запущенных на машине жертвы, и выявил среди них процессы десктопных LLM-приложений.</p></a>
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0008 Выявление</span><p>Исследователи обнаруживают, что ИИ-агент имеет доступ к источнику данных «Customer Support Account Owners.csv».</p></a>
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0008 Выявление</span><p>Исследователи делают вывод, что у ИИ-агента есть инструмент для отправки писем.</p></a>
</div>


Показано 12 из 19 примеров.
