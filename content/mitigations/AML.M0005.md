---
atlas_id: AML.M0005
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Policy
created_date: "2023-04-12"
description: Настройте контроль доступа к внутренним реестрам моделей и ограничьте внутренний доступ к моделям, используемым в продакшене. Ограничьте доступ к обучающим данным только для авторизованных пользователей.
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Business and Data Understanding
    - Data Preparation
    - AI Model Engineering
    - AI Model Evaluation
modified_date: "2026-07-31"
source_name: Control Access to AI Models and Data at Rest
technique_count: 20
title: Контроль доступа к ИИ-моделям и хранимым данным
url: /mitigations/AML.M0005/
---

Настройте контроль доступа к внутренним реестрам моделей и ограничьте внутренний доступ к моделям, используемым в продакшене.
Ограничьте доступ к обучающим данным только для авторизованных пользователей.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0007/"><span class="relation-id">AML.T0007</span><strong>Выявление ИИ-артефактов</strong><p>Контроль доступа может ограничить способность злоумышленника выявлять ИИ-модели, наборы данных и другие артефакты в системе.</p></a>
<a class="relation-item" href="/techniques/AML.T0010.002/"><span class="relation-id">AML.T0010.002</span><strong>Данные</strong><p>Контроль доступа может предотвращать подмену ML-артефактов и их неавторизованное копирование.</p></a>
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><p>Контроль доступа может предотвращать подмену ML-артефактов и их неавторизованное копирование.</p></a>
<a class="relation-item" href="/techniques/AML.T0012/"><span class="relation-id">AML.T0012</span><strong>Действующие учетные записи</strong><p>Предоставляйте доступ к реестрам моделей и обучающим данным только одобренным субъектам доступа, чтобы при компрометации учётной записи доступными оказывались лишь те ИИ-активы, доступ к которым для этой учётной записи был явно разрешён.</p></a>
<a class="relation-item" href="/techniques/AML.T0018/"><span class="relation-id">AML.T0018</span><strong>Манипуляция ИИ-моделью</strong><p>Контроль доступа может предотвращать подмену ИИ-артефактов и их неавторизованное изменение.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.000/"><span class="relation-id">AML.T0018.000</span><strong>Отравление ИИ-модели</strong><p>Контроль доступа может предотвращать подмену ML-артефактов и их неавторизованное копирование.</p></a>
<a class="relation-item" href="/techniques/AML.T0018.001/"><span class="relation-id">AML.T0018.001</span><strong>Изменение архитектуры ИИ-модели</strong><p>Контроль доступа может предотвращать подмену ML-артефактов и их неавторизованное копирование.</p></a>
<a class="relation-item" href="/techniques/AML.T0020/"><span class="relation-id">AML.T0020</span><strong>Отравление обучающих данных</strong><p>Контроль доступа может предотвращать подмену ML-артефактов и их неавторизованное копирование.</p></a>
<a class="relation-item" href="/techniques/AML.T0021/"><span class="relation-id">AML.T0021</span><strong>Создание учетных записей</strong><p>Проверяйте идентичность субъекта перед предоставлением доступа к реестрам моделей, чтобы вновь созданные учётные записи не получали автоматического доступа к защищённым ИИ-артефактам.</p></a>
<a class="relation-item" href="/techniques/AML.T0025/"><span class="relation-id">AML.T0025</span><strong>Эксфильтрация киберсредствами</strong><p>Контроль доступа может предотвращать эксфильтрацию.</p></a>
<a class="relation-item" href="/techniques/AML.T0035/"><span class="relation-id">AML.T0035</span><strong>Сбор ИИ-артефактов</strong><p>Контроль доступа может предотвращать или ограничивать сбор ИИ-артефактов в системе жертвы.</p></a>
<a class="relation-item" href="/techniques/AML.T0042/"><span class="relation-id">AML.T0042</span><strong>Проверка атаки</strong><p>Контроль доступа к моделям в состоянии покоя может помешать злоумышленнику проверять эффективность атаки.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.000/"><span class="relation-id">AML.T0043.000</span><strong>Оптимизация в режиме белого ящика</strong><p>Контроль доступа может сократить избыточный доступ к ИИ-моделям и не дать злоумышленнику получить доступ в режиме белого ящика.</p></a>
<a class="relation-item" href="/techniques/AML.T0044/"><span class="relation-id">AML.T0044</span><strong>Полный доступ к ИИ-модели</strong><p>Контроль доступа к моделям и данным в состоянии покоя может помочь предотвратить полный доступ к модели.</p></a>
<a class="relation-item" href="/techniques/AML.T0048.004/"><span class="relation-id">AML.T0048.004</span><strong>Кража интеллектуальной собственности ИИ</strong><p>Контроль доступа может предотвращать кражу интеллектуальной собственности.</p></a>
<a class="relation-item" href="/techniques/AML.T0069/"><span class="relation-id">AML.T0069</span><strong>Выявление системной информации LLM</strong><p>Ограничивайте доступ к хранящимся системным промптам, конфигурационным файлам и метаданным моделей.</p></a>
<a class="relation-item" href="/techniques/AML.T0069.000/"><span class="relation-id">AML.T0069.000</span><strong>Наборы специальных символов</strong><p>Ограничивайте доступ к хранящимся шаблонам промптов и конфигурациям, содержащим внутренние разделители.</p></a>
<a class="relation-item" href="/techniques/AML.T0069.001/"><span class="relation-id">AML.T0069.001</span><strong>Ключевые слова системных инструкций</strong><p>Ограничивайте доступ к хранящимся системным инструкциям и определениям инструментов, содержащим привилегированные ключевые слова.</p></a>
<a class="relation-item" href="/techniques/AML.T0069.002/"><span class="relation-id">AML.T0069.002</span><strong>Системный промпт</strong><p>Ограничивайте доступ к хранящимся системным промптам и шаблонам промптов.</p></a>
<a class="relation-item" href="/techniques/AML.T0112.001/"><span class="relation-id">AML.T0112.001</span><strong>ИИ-артефакты</strong><p>Ограничивайте права на запись в реестры моделей и хранилища ИИ-артефактов.</p></a>
</div>
