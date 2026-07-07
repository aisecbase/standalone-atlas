---
atlas_id: AML.T0052.001
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2026-04-22"
description: Злоумышленники могут использовать дипфейки, то есть синтетические изображения, аудио или видео, созданные с помощью ИИ, в фишинговых кампаниях, чтобы выдавать себя за доверенных лиц, руководителей или организации....
generated: true
generated_by: atlasgen
maturity: feasible
mitigation_count: 2
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 0
source_name: Deepfake-Assisted Phishing
subtechnique_count: 0
subtechnique_of: AML.T0052
tactics:
    - AML.TA0004
    - AML.TA0015
title: Фишинг с использованием дипфейков
url: /techniques/AML.T0052.001/
---

Злоумышленники могут использовать дипфейки, то есть синтетические изображения, аудио или видео, созданные с помощью ИИ, в фишинговых кампаниях, чтобы выдавать себя за доверенных лиц, руководителей или организации. Такие атаки эксплуатируют доверие людей: мошеннические голосовые или видеосообщения подаются как легитимные, что позволяет злоумышленникам манипулировать целями и добиваться раскрытия учетных данных, перевода средств или предоставления доступа к системам.

Голосовые дипфейки, то есть голоса, клонированные с помощью ИИ, используются в атаках вишинга [\[1\]][vishing] (голосового фишинга) по телефону или VoIP. Злоумышленники могут клонировать голос цели, используя несколько секунд [\[2\]][valle] публично доступного аудио из выступлений, звонков с отчетностью, подкастов или социальных сетей [\[3\]][voice]. Затем такие клонированные голоса используются в заранее записанных голосовых сообщениях или живых телефонных звонках. Видеодипфейки могут имитировать лицо и голос доверенного человека. Злоумышленники используют публично доступные видео с корпоративных встреч, звонков с отчетностью или из социальных сетей, чтобы создавать убедительные ИИ-сгенерированные видео целевых лиц. Такие видео используются в живых видеоконференциях или записанных видеосообщениях. ИИ-сгенерированный контент развился до уровня, при котором его часто сложно распознать как синтетический [\[4\]][fbi].

При подготовке к [фишинговой](/techniques/AML.T0052) кампании злоумышленники могут сначала выполнить [получение возможностей](/techniques/AML.T0016): [генеративный ИИ](/techniques/AML.T0016.002), а затем [создание дипфейков](/techniques/AML.T0088). Фишинговые кампании с дипфейками часто используют и другие каналы коммуникации, такие как email, SMS или мессенджеры, для многоуровневых атак социальной инженерии [\[5\]][aiid839].

Такие атаки охватывают широкий спектр жертв и типов атак, демонстрируя масштаб мошенничества с использованием дипфейков. Злоумышленники проводили масштабные фишинговые кампании с дипфейками против отдельных людей, включая целевое мошенничество [\[6\]][aiid564] [\[7\]][oecd1] [\[8\]][aiid1280] [\[9\]][aiid1285], а также крупномасштабные кампании по сбору учетных данных, нацеленные на миллиарды пользователей [\[10\]][aiid839] [\[11\]][aiid941]. Злоумышленники использовали дипфейки для имперсонации руководителей [\[12\]][aiid1100], что приводило к значительным финансовым потерям для компаний [\[13\]][aiid634] [\[14\]][aiid147]. Также есть сообщения о массовых кампаниях, в которых целями становились государственные должностные лица [\[4\]][fbi] [\[15\]][aiid927].

Эти атаки охватывают разные каналы коммуникации, включая голосовые дипфейки для вишинга [\[16\]][aiid567] и видеодипфейки в конференц-звонках [\[13\]][aiid634], а также многоканальные кампании, совмещающие телефон, email и платформы обмена сообщениями [\[10\]][aiid839].

[valle]: https://www.microsoft.com/en-us/research/project/vall-e-x/ "VALL-E Family: Neural codec language models for speech synthesis"

[vishing]: https://www.social-engineer.org/framework/attack-vectors/vishing/ "Vishing - Social-Engineer Framework"

[voice]: https://cloud.google.com/blog/topics/threat-intelligence/ai-powered-voice-spoofing-vishing-attacks "AI-powered voice spoofing: Understanding and defending against vishing attacks"

[fbi]: https://www.ic3.gov/PSA/2025/PSA250515/ "FBI Public Service Advisory: Scammers are deepfaking voices of senior US government officials"

[oecd1]: https://oecd.ai/en/incidents/2026-04-06-ca7a "AI-Generated Voice Used in Scam Targeting Drica Moraes' Contacts"

[aiid634]: https://incidentdatabase.ai/cite/634/ "Alleged Deepfake CFO Scam Reportedly Costs Multinational Engineering Firm Arup $25 Million"

[aiid147]: https://incidentdatabase.ai/cite/147/ "Reported AI-Cloned Voice Used to Deceive Hong Kong Bank Manager in Purported $35 Million Fraud Scheme"

[aiid1100]: https://incidentdatabase.ai/cite/1100/ "AI Incident Database - LastPass CEO Voice Deepfake Attempt"

[aiid927]: https://incidentdatabase.ai/cite/927/ "Italian Defense Minister Voice Clone"

[aiid564]: https://incidentdatabase.ai/cite/564/ "Voice deepfake targets bank in failed transfer scam"

[aiid567]: https://incidentdatabase.ai/cite/567/ "Deepfake Voice Exploit Compromises Retool's Cloud Services"

[aiid1280]: https://incidentdatabase.ai/cite/1280/ "Reported Use of AI Voice and Identity Manipulation in the 'Phantom Hacker' Fraud Scheme"

[aiid1285]: https://incidentdatabase.ai/cite/1285/ "Purportedly AI-Generated Jason Momoa Deepfake Used in Romance Scam"

[aiid839]: https://incidentdatabase.ai/cite/839/ "Purportedly AI-Driven Phishing Scam Uses Spoofed Google Call to Attempt Gmail Breach"

[aiid941]: https://incidentdatabase.ai/cite/941/ "AI-Driven Phishing Scam Uses Deepfake Robocalls to Target Gmail Users"


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0004/"><span class="relation-id">AML.TA0004</span><strong>Первичный доступ</strong></a>
<a class="relation-item" href="/tactics/AML.TA0015/"><span class="relation-id">AML.TA0015</span><strong>Латеральное перемещение</strong></a>
</div>


## Родительская техника

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0052/"><span class="relation-id">AML.T0052</span><strong>Фишинг</strong></a>
</div>


## Другие подтехники родителя

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0052.000/"><span class="relation-id">AML.T0052.000</span><strong>Целевой фишинг через LLM для социальной инженерии</strong><span class="relation-meta">Подтехника</span></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0018/"><span class="relation-id">AML.M0018</span><strong>Обучение пользователей</strong><p>Обучайте пользователей угрозам, связанным с дипфейками, в том числе распознаванию синтетического голоса, видео и текста. Рекомендуйте проверку по независимым каналам, например по известному номеру для обратного звонка, перед обработкой чувствительных запросов или предоставлением чувствительной информации по голосовой или видеосвязи.</p></a>
<a class="relation-item" href="/mitigations/AML.M0034/"><span class="relation-id">AML.M0034</span><strong>Обнаружение дипфейков</strong><p>Разверните технические средства контроля для обнаружения и блокировки синтетического аудио и видео. К ним относятся инструменты анализа на основе ИИ, которые проверяют медиа на артефакты, указывающие на дипфейки.</p></a>
</div>


## Источники

- [AI Incident Database - LastPass CEO Voice Deepfake Attempt](https://incidentdatabase.ai/cite/1100/)
- [Reported Use of AI Voice and Identity Manipulation in the 'Phantom Hacker' Fraud Scheme](https://incidentdatabase.ai/cite/1280/)
- [Purportedly AI-Generated Jason Momoa Deepfake Used in Romance Scam](https://incidentdatabase.ai/cite/1285/)
- [Reported AI-Cloned Voice Used to Deceive Hong Kong Bank Manager in Purported $35 Million Fraud Scheme](https://incidentdatabase.ai/cite/147/)
- [Voice deepfake targets bank in failed transfer scam](https://incidentdatabase.ai/cite/564/)
- [Deepfake Voice Exploit Compromises Retool's Cloud Services](https://incidentdatabase.ai/cite/567/)
- [Alleged Deepfake CFO Scam Reportedly Costs Multinational Engineering Firm Arup $25 Million](https://incidentdatabase.ai/cite/634/)
- [Purportedly AI-Driven Phishing Scam Uses Spoofed Google Call to Attempt Gmail Breach](https://incidentdatabase.ai/cite/839/)
- [Italian Defense Minister Voice Clone](https://incidentdatabase.ai/cite/927/)
- [AI-Driven Phishing Scam Uses Deepfake Robocalls to Target Gmail Users](https://incidentdatabase.ai/cite/941/)
- [FBI Public Service Advisory: Scammers are deepfaking voices of senior US government officials](https://www.ic3.gov/PSA/2025/PSA250515/)
- [AI-Generated Voice Used in Scam Targeting Drica Moraes' Contacts](https://oecd.ai/en/incidents/2026-04-06-ca7a)
- [VALL-E Family: Neural codec language models for speech synthesis](https://www.microsoft.com/en-us/research/project/vall-e-x/)
- [Vishing - Social-Engineer Framework](https://www.social-engineer.org/framework/attack-vectors/vishing/)
- [AI-powered voice spoofing: Understanding and defending against vishing attacks](https://cloud.google.com/blog/topics/threat-intelligence/ai-powered-voice-spoofing-vishing-attacks)
