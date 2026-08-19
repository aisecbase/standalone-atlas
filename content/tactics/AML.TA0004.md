---
atlas_id: AML.TA0004
atlas_type: tactic
attack_ref_id: TA0001
attack_ref_url: https://attack.mitre.org/tactics/TA0001/
created_date: "2022-01-24"
description: Злоумышленник пытается получить доступ к ИИ-системе. Целевой системой может быть сеть, мобильное устройство или периферийное устройство, например сенсорная платформа. ИИ-возможности, используемые системой, могут быть...
generated: true
generated_by: atlasgen
modified_date: "2025-04-09"
procedure_count: 53
source_name: Initial Access
technique_count: 23
title: Первичный доступ
url: /tactics/AML.TA0004/
---

Злоумышленник пытается получить доступ к ИИ-системе.

Целевой системой может быть сеть, мобильное устройство или периферийное устройство, например сенсорная платформа.
ИИ-возможности, используемые системой, могут быть локальными, встроенными в устройство, или облачными.

Первичный доступ включает техники, которые используют различные векторы входа для первоначального закрепления в системе.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0010/"><span class="relation-id">AML.T0010</span><strong>Компрометация цепочки поставок ИИ</strong></a>
<a class="relation-item" href="/techniques/AML.T0010.000/"><span class="relation-id">AML.T0010.000</span><strong>Аппаратное обеспечение</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.000/"><span class="relation-id">AML.T0010.000</span><strong>Аппаратное обеспечение</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.001/"><span class="relation-id">AML.T0010.001</span><strong>ПО для ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.001/"><span class="relation-id">AML.T0010.001</span><strong>ПО для ИИ</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.002/"><span class="relation-id">AML.T0010.002</span><strong>Данные</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.002/"><span class="relation-id">AML.T0010.002</span><strong>Данные</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.003/"><span class="relation-id">AML.T0010.003</span><strong>Модель</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.004/"><span class="relation-id">AML.T0010.004</span><strong>Реестр контейнеров</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.004/"><span class="relation-id">AML.T0010.004</span><strong>Реестр контейнеров</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.005/"><span class="relation-id">AML.T0010.005</span><strong>Инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0010.005/"><span class="relation-id">AML.T0010.005</span><strong>Инструмент ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0012/"><span class="relation-id">AML.T0012</span><strong>Действующие учетные записи</strong></a>
<a class="relation-item" href="/techniques/AML.T0015/"><span class="relation-id">AML.T0015</span><strong>Обход ИИ-модели</strong></a>
<a class="relation-item" href="/techniques/AML.T0049/"><span class="relation-id">AML.T0049</span><strong>Эксплуатация приложения, доступного из интернета</strong></a>
<a class="relation-item" href="/techniques/AML.T0052/"><span class="relation-id">AML.T0052</span><strong>Фишинг</strong></a>
<a class="relation-item" href="/techniques/AML.T0052.000/"><span class="relation-id">AML.T0052.000</span><strong>Целевой фишинг через LLM для социальной инженерии</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0052.000/"><span class="relation-id">AML.T0052.000</span><strong>Целевой фишинг через LLM для социальной инженерии</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0052.001/"><span class="relation-id">AML.T0052.001</span><strong>Фишинг с использованием дипфейков</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0052.001/"><span class="relation-id">AML.T0052.001</span><strong>Фишинг с использованием дипфейков</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0078/"><span class="relation-id">AML.T0078</span><strong>Компрометация при посещении сайта</strong></a>
<a class="relation-item" href="/techniques/AML.T0093/"><span class="relation-id">AML.T0093</span><strong>Внедрение промпта через публичное приложение</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0002/"><span class="relation-id">AML.CS0002</span><strong>Отравление VirusTotal</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0004 Первичный доступ</span><p>Злоумышленник загрузил мутированные образцы на платформу.</p></a>
<a class="relation-item" href="/studies/AML.CS0004/"><span class="relation-id">AML.CS0004</span><strong>Атака на систему распознавания лиц через подмену видеопотока камеры</strong><span class="relation-meta">Актор: Two individuals / Тактика: AML.TA0004 Первичный доступ</span><p>Злоумышленники успешно обошли систему распознавания лиц. Это позволило им выдать себя за жертву и подтвердить ее личность в налоговой системе.</p></a>
<a class="relation-item" href="/studies/AML.CS0009/"><span class="relation-id">AML.CS0009</span><strong>Отравление Tay</strong><span class="relation-meta">Актор: 4chan Users / Тактика: AML.TA0004 Первичный доступ</span><p>Бот Tay использовал взаимодействия с пользователями Twitter как обучающие данные, чтобы улучшать свои диалоги. Злоумышленники смогли скоординироваться и использовать эту петлю обратной связи, чтобы исказить поведение Tay.</p></a>
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0004 Первичный доступ</span><p>Команда использовала действительную учетную запись, чтобы получить доступ к сети.</p></a>
<a class="relation-item" href="/studies/AML.CS0012/"><span class="relation-id">AML.CS0012</span><strong>Обход системы идентификации лиц с помощью физических контрмер</strong><span class="relation-meta">Актор: MITRE AI Red Team / Тактика: AML.TA0004 Первичный доступ</span><p>Команда получила доступ к коммерческому сервису идентификации лиц и его API через действительную учетную запись.</p></a>
<a class="relation-item" href="/studies/AML.CS0013/"><span class="relation-id">AML.CS0013</span><strong>Бэкдор-атака на модели глубокого обучения в мобильных приложениях</strong><span class="relation-meta">Актор: Yuanchun Li, Jiayi Hua, Haoyu Wang, Chunyang Chen, Yunxin Liu / Тактика: AML.TA0004 Первичный доступ</span><p>На практике вредоносный APK-файл должен быть установлен на устройства жертв через компрометацию цепочки поставок.</p></a>
<a class="relation-item" href="/studies/AML.CS0015/"><span class="relation-id">AML.CS0015</span><strong>Компрометация цепочки зависимостей PyTorch</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0004 Первичный доступ</span><p>Вредоносный пакет зависимости с именем `torchtriton` был загружен в репозиторий PyPI с тем же именем, что и пакет, поставлявшийся со сборкой PyTorch-nightly. Этот вредоносный пакет содержал дополнительный код, отправлявший чувствительные данные с машины. Вредоносный `torchtriton` устанавливался вместо легитимного пакета, потому что PyPI имел приоритет над другими источниками. Подробнее см. [этот GitHub issue](https://github.com/pypa/pip/issues/8606).</p></a>
<a class="relation-item" href="/studies/AML.CS0016/"><span class="relation-id">AML.CS0016</span><strong>Выполнение кода в MathGPT через промпт-инъекцию</strong><span class="relation-meta">Актор: Ludwig-Ferdinand Stumpp / Тактика: AML.TA0004 Первичный доступ</span><p>Это показало, что исследователь может использовать уязвимость к промпт-инъекции в модели GPT-3, применяемой в MathGPT, как вектор первичного доступа.</p></a>
<a class="relation-item" href="/studies/AML.CS0017/"><span class="relation-id">AML.CS0017</span><strong>Обход проверки личности ID.me</strong><span class="relation-meta">Актор: One individual / Тактика: AML.TA0004 Первичный доступ</span><p>Мужчина собрал украденные персональные данные, включая имена, даты рождения и номера Social Security, и использовал их вместе со своими фотографиями в париках для получения поддельных водительских удостоверений. Он загружал поддельные документы вместе с селфи. Система проверки документов ID.me сопоставляла селфи с фотографией в документе, что позволяло части мошеннических заявок продвигаться дальше по процессу обработки.</p></a>
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0004 Первичный доступ</span><p>Jupyter Notebook часто используются для исследований и экспериментов в области ML и data science и содержат исполняемые фрагменты Python-кода, а также типовую функциональность командной строки Unix. Пользователи могут столкнуться со скомпрометированным notebook на публичных сайтах или получить его напрямую по ссылке.</p></a>
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0004 Первичный доступ</span><p>Пользователь-жертва может подключить свой Google Drive к скомпрометированному Colab notebook. Типовые причины подключения ML-notebook к Google Drive включают обучение на размещенных там данных или сохранение выходных файлов модели. При выполнении появляется окно подтверждения доступа и предупреждение о возможном доступе к данным: &gt; Этот notebook запрашивает доступ к вашим файлам Google Drive. Предоставление доступа к Google Drive позволит коду, выполняемому в notebook, изменять файлы в вашем Google Drive. Перед предоставлением доступа обязательно проверьте код notebook. Тем не менее пользователь-жертва может принять запрос и предоставить скомпрометированному Colab notebook доступ к своему Drive. Выданные разрешения включают: - создание, изменение и удаление всех файлов Google Drive; - просмотр данных Google Photos; - просмотр контактов Google.</p></a>
<a class="relation-item" href="/studies/AML.CS0019/"><span class="relation-id">AML.CS0019</span><strong>PoisonGPT</strong><span class="relation-meta">Актор: Mithril Security Researchers / Тактика: AML.TA0004 Первичный доступ</span><p>Ничего не подозревающие пользователи могли скачать состязательную модель и интегрировать ее в приложения. После раскрытия информации об упражнении Hugging Face отключил репозиторий с похожим именем.</p></a>
</div>


Показано 12 из 53 примеров.
