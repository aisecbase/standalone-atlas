---
atlas_id: AML.M0003
atlas_type: mitigation
attack_ref_id: ""
attack_ref_url: ""
category:
    - Technical - AI
created_date: "2023-04-12"
description: Проектируйте и обучайте модели предиктивного ИИ так, чтобы они сохраняли заданные показатели качества при обработке состязательных примеров. К состязательным примерам могут относиться входные данные с внесёнными...
generated: true
generated_by: atlasgen
ml_lifecycle:
    - Data Preparation
    - AI Model Engineering
modified_date: "2026-07-31"
source_name: Predictive AI Model Hardening
technique_count: 8
title: Повышение устойчивости моделей предиктивного ИИ
url: /mitigations/AML.M0003/
---

Проектируйте и обучайте модели предиктивного ИИ так, чтобы они сохраняли заданные показатели качества при обработке состязательных примеров. К состязательным примерам могут относиться входные данные с внесёнными цифровыми возмущениями или физические средства противодействия модели, предназначенные для того, чтобы вызвать ошибочную классификацию, невыявление объекта либо иной выбранный злоумышленником результат прогнозирования.

К методам обеспечения устойчивости могут относиться состязательное обучение, устойчивые архитектуры моделей, защитная дистилляция и методы сертифицированной устойчивости.


## Связанные техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0015/"><span class="relation-id">AML.T0015</span><strong>Обход ИИ-модели</strong><p>Модели с повышенной устойчивостью труднее обойти.</p></a>
<a class="relation-item" href="/techniques/AML.T0031/"><span class="relation-id">AML.T0031</span><strong>Нарушение целостности ИИ-модели</strong><p>Модели с повышенной устойчивостью менее подвержены атакам на целостность.</p></a>
<a class="relation-item" href="/techniques/AML.T0043/"><span class="relation-id">AML.T0043</span><strong>Создание состязательных данных</strong><p>Модели с повышенной устойчивостью лучше противостоят состязательным входным данным.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.000/"><span class="relation-id">AML.T0043.000</span><strong>Оптимизация в режиме белого ящика</strong><p>Модели с повышенной устойчивостью лучше противостоят состязательным входным данным.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.001/"><span class="relation-id">AML.T0043.001</span><strong>Оптимизация в режиме чёрного ящика</strong><p>Модели с повышенной устойчивостью лучше противостоят состязательным входным данным.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.002/"><span class="relation-id">AML.T0043.002</span><strong>Перенос на модель чёрного ящика</strong><p>Модели с повышенной устойчивостью лучше противостоят состязательным входным данным.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.003/"><span class="relation-id">AML.T0043.003</span><strong>Ручная модификация</strong><p>Модели с повышенной устойчивостью лучше противостоят состязательным входным данным.</p></a>
<a class="relation-item" href="/techniques/AML.T0043.004/"><span class="relation-id">AML.T0043.004</span><strong>Добавление бэкдор-триггера</strong><p>Модели с повышенной устойчивостью лучше противостоят состязательным входным данным.</p></a>
</div>
