package trans

import "github.com/strongo/bots-framework/core"

const adsCommandTitle = "\xE2\xAD\x90\xE2\xAD\x90\xE2\xAD\x90"

// TRANS - translation string
var TRANS = map[string]map[string]string{
	"EXAMPLE": {
		"en-US": "SAMPLE",
		"es-ES": "EJEMPLO",
		"fa-IR": "نمونه",
		"it-IT": "ESEMPIO",
		"ru-RU": "ПРИМЕР",
	},

	"Jan": {
		"en-US": "Jan",
		"es-ES": "En",
		"fa-IR": "ژانویه",
		"it-IT": "Gen",
		"ru-RU": "Янв.",
	},
	"Feb": {
		"en-US": "Feb",
		"es-ES": "Feb",
		"fa-IR": "فوریه",
		"it-IT": "Feb",
		"ru-RU": "Фев.",
	},
	"Mar": {
		"en-US": "Mar",
		"es-ES": "Mar",
		"fa-IR": "مارس",
		"it-IT": "Mar",
		"ru-RU": "Мрт.",
	},
	"Apr": {
		"en-US": "Apr",
		"es-ES": "Abr",
		"fa-IR": "آوریل",
		"it-IT": "Apr",
		"ru-RU": "Апр.",
	},
	"May": {
		"en-US": "May",
		"es-ES": "May",
		"fa-IR": "مه",
		"it-IT": "Mag",
		"ru-RU": "Май ",
	},
	"Jun": {
		"en-US": "Jun",
		"es-ES": "Jun",
		"fa-IR": "ژوئن",
		"it-IT": "Giu",
		"ru-RU": "Июнь",
	},
	"Jul": {
		"en-US": "Jul",
		"es-ES": "Jul",
		"fa-IR": "ژوئیه",
		"it-IT": "Lug",
		"ru-RU": "Июль",
	},
	"Aug": {
		"en-US": "Aug",
		"es-ES": "Ago",
		"fa-IR": "اوت",
		"it-IT": "Ago",
		"ru-RU": "Авг.",
	},
	"Sep": {
		"en-US": "Sep",
		"es-ES": "Sep",
		"fa-IR": "سپتامبر",
		"it-IT": "Sett",
		"ru-RU": "Сен.",
	},
	"Oct": {
		"en-US": "Oct",
		"es-ES": "Oct",
		"fa-IR": "اکتبر",
		"it-IT": "Ott",
		"ru-RU": "Окт.",
	},
	"Nov": {
		"en-US": "Nov",
		"es-ES": "Nov",
		"fa-IR": "نوامبر",
		"it-IT": "Nov",
		"ru-RU": "Нбр.",
	},
	"Dec": {
		"en-US": "Dec",
		"es-ES": "Dic",
		"fa-IR": "دسامبر",
		"it-IT": "Dic",
		"ru-RU": "Дек.",
	},
	COMMAND_START: {
		"en-US": "start",
		"es-ES": "inicio",
		"fa-IR": "شروع",
		"it-IT": "inizio",
		"ru-RU": "старт",
	},
	COMMAND_MENU: {
		"en-US": "menu",
		"es-ES": "menu",
		"fa-IR": "منو",
		"it-IT": "menu", // TODO(IT): Google translated
		"ru-RU": "меню",
	},
	COMMAND_GAVE: {
		"en-US": "gave",
		"es-ES": "débito",
		"fa-IR": "قرض_دادن",
		"it-IT": "debito",
		"ru-RU": "дал",
	},
	COMMAND_GOT: {
		"en-US": "got",
		"es-ES": "crédito",
		"fa-IR": "قرض_گرفتن",
		"it-IT": "credito",
		"ru-RU": "взял",
	},
	COMMAND_RETURNED: {
		"en-US": "returned",
		"es-ES": "devuelto",
		"fa-IR": "بازگردانده_شده",
		"it-IT": "rientra",
		"ru-RU": "вернул",
	},
	COMMAND_BALANCE: {
		"en-US": "balance",
		"es-ES": "balance",
		"fa-IR": "تراز",
		"it-IT": "bilancio",
		"ru-RU": "баланс",
	},
	COMMAND_HISTORY: {
		"en-US": "history",
		"es-ES": "cronología",
		"fa-IR": "سوابق",
		"it-IT": "cronologia",
		"ru-RU": "история",
	},
	COMMAND_SETTINGS: {
		"en-US": "settings",
		"es-ES": "ajustes",
		"fa-IR": "تنظیمات",
		"it-IT": "impostazioni",
		"ru-RU": "настройки",
	},
	COMMAND_HELP: {
		"en-US": "help",
		"es-ES": "ayuda",
		"fa-IR": "کمک",
		"it-IT": "aiuto",
		"ru-RU": "помощь",
	},
	COMMAND_CANCEL: {
		"en-US": "cancel",
		"es-ES": "cancelar",
		"fa-IR": "کنسل",
		"it-IT": "annulla",
		"ru-RU": "/отменить",
	},
	COMMAND_CLEAR: {
		"en-US": "clear",
		"es-ES": "borrar",
		"fa-IR": "پاک_کردن",
		"it-IT": "chiaro",
		"ru-RU": "очистить",
	},
	adsCommandTitle: {
		"en-US": adsCommandTitle,
		"es-ES": adsCommandTitle,
		"fa-IR": adsCommandTitle,
		"it-IT": adsCommandTitle,
		"ru-RU": adsCommandTitle,
	},
	" and ": {
		"en-US": " and ",
		"es-ES": " y ",
		"fa-IR": " و ",
		"it-IT": " e ",
		"ru-RU": " и ",
	},
	bots.MESSAGE_TEXT_OOPS_SOMETHING_WENT_WRONG: {
		"en-US": "Oops, something went wrong... \xF0\x9F\x98\xB3",
		"es-ES": "Ops,  algo ha salido mal... \xF0\x9F\x98\xB3",
		"fa-IR": "اوه، یک جای کار مشکل دارد ...  \xF0\x9F\x98\xB3",
		"it-IT": "Ops, qualcosa e' andato storto... \xF0\x9F\x98\xB3",
		"ru-RU": "Упс, что-то пошло не так... \xF0\x9F\x98\xB3",
	},
	MESSAGE_TEXT_ASK_DUE: {
		"en-US": "When is the due date?",
		"es-ES": "¿Cuándo es la fecha de devolución?",
		"fa-IR": "سررسید چه زمانی است؟",
		"it-IT": "Data di scadenza?",
		"ru-RU": "Когда дата возврата?",
	},
	MESSAGE_TEXT_ASK_DATE_TO_REMIND: {
		"en-US": `To set date for next reminder please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,

		"es-ES": `Para establecer la fecha recordatoria escribela en el siguiente formato <i>DD.MM.AÑO</i>.
<b>For example</b> para 20 de Enero 2017 envia:
    <i>20.01.2017</i>`,
		
		"fa-IR": `لطفاً برای تنظیم یادآور بعدی آنرا با متنی با این فرمت ارسال نمایید. <i>روز.ماه.سال</i>.
<b>برای مثال</b> برای 15 خرداد 1396 ثبت کنید:
    <i>15.03.1396</i>`,

		"it-IT": `Per impostare la data per il promemoria successivo invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
		<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:

		<i>20.01.2017</i>`,
		
		"ru-RU": `Чтобы задать дату напопинания напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г. отправьте:
    <i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_ASK_DUE_DATE: {
		"en-US": `To set due date please send it as a text in format of <i>DD.MM.YEAR</i>.
		<b>For example</b> for 20th of January 2017 submit:
		<i>20.01.2017</i>`,

		"es-ES": `Para establecer la fecha de devolución escribela en el siguiente formato <i>DD.MM.AÑO</i>.
                 <b>For example</b> para 20 de Enero 2017 envia:
                 <i>20.01.2017</i>`,

		"fa-IR": `لطفاً برای تنظیم تاریخ سررسید این فرمت را رعایت فرمایید.<i>روز.ماه.سال</i>.
		<b>برای مثال</b> برای 20 ژانویه 2017 ثبت کنید:
		<i>20.01.2017</i>`,

		"it-IT": `Per impostare la data di scadenza invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
		<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:
		<i>20.01.2017</i>`,
		
		"ru-RU": `Чтобы задать дату возврата напишите её в формате <i>ДД.MM.ГОД</i>.
		<b>Например</b> для 20 января 2017 г.отправьте:
		<i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_WRONG_DATE: {
		"en-US": "Sorry, there is something wrong with the date you've provided.",
		"es-ES": "Lo siento, algo no es correcto con la fecha que has puesto",
		"fa-IR": "متاسفم، در تاریخی که وارد نمودید مشکلی وجود دارد.",
		"it-IT": "Mi spiace, ma c'e' qualcosa di sbagliato nella data che hai inserito.",
		"ru-RU": "Извините, что-то не так с датой которую вы отправили.",
	},
	COMMAND_TEXT_DISABLE_REMINDER: {
		"en-US": "No reminder",
		"es-ES": "No recordar",
		"fa-IR": "بدون یادآور",
		"it-IT": "Nessun promemoria",
		"ru-RU": "Не напоминать",
	},
	COMMAND_TEXT_TOMORROW: {
		"en-US": "Tomorrow",
		"es-ES": "Mañana",
		"fa-IR": "فردا",
		"it-IT": "Domani",
		"ru-RU": "Завтра",
	},
	COMMAND_TEXT_DAY_AFTER_TOMORROW: {
		"en-US": "Day after tomorrow",
		"es-ES": "Pasada mañana",
		"fa-IR": "پس فردا",
		"it-IT": "Dopo domani",
		"ru-RU": "Послезавтра",
	},
	COMMAND_TEXT_THIS_WEEK: {
		"en-US": "This week",
		"es-ES": "Esta semana",
		"fa-IR": "این هفته",
		"it-IT": "Questa settimana",
		"ru-RU": "На этой неделе",
	},
	COMMAND_TEXT_YES_IT_HAS_RETURN_DEADLINE: {
		"en-US": "Yes, it has a deadline!",
		"es-ES": "Si, hay una fecha de devolución!",
		"fa-IR": "بله، دارای آخرین فرصت می باشد!",
		"it-IT": "Si, c'e' una data di scadenza",
		"ru-RU": "Да, есть срок возврата!",
	},
	COMMAND_TEXT_NO_IT_CAN_BE_RETURNED_ANYTIME: {
		"en-US": "No, whenever is fine.",
		"es-ES": "No, sin fecha límite.",
		"fa-IR": "خیر، هر زمانی مناسب است.",
		"it-IT": "No, nessuna scadenza",
		"ru-RU": "Нет, срок возврата не важен.",
	},
	COMMAND_TEXT_IT_CAN_BE_RETURNED_ANYTIME: {
		"en-US": "Whenever is fine",
		"es-ES": "Cualquier día",
		"fa-IR": "هر زمانی مناسب است.",
		"it-IT": "No, Nessuna scadenza",
		"ru-RU": "Когда-нибудь",
	},
	COMMAND_TEXT_IN_FEW_MINUTES: {
		"en-US": "In few minutes",
		"es-ES2: "En unos minutos"
		"fa-IR": "در چند دقیقه",
		"it-IT": "Fra alcuni minuti",
		"ru-RU": "Через минуту",
	},
	COMMAND_TEXT_IN_1_WEEK: {
		"en-US": "In 1 week",
		"es-ES": "En una semana",
		"fa-IR": "ظرف یک هفته",
		"it-IT": "Fra una settimana",
		"ru-RU": "Через неделю",
	},
	COMMAND_TEXT_IN_1_MONTH: {
		"en-US": "In 1 month",
		"es-ES": "En un mes",
		"fa-IR": "ظرف یک ماه",
		"it-IT": "Fra un mese",
		"ru-RU": "Через месяц",
	},
	COMMAND_TEXT_SET_DATE: {
		"en-US": "Set date",
		"es-ES": "Establecer la fecha",
		"fa-IR": "تاریخ را تنظیم کنید",
		"it-IT": "Imposta la data",
		"ru-RU": "Задать дату",
	},
	COMMAND_TEXT_MONDAY: {
		"en-US": "Monday",
		"es-ES": "Lunes",
		"fa-IR": "دوشنبه",
		"it-IT": "Lunedi'",
		"ru-RU": "Понедельник",
	},
	COMMAND_TEXT_TUESDAY: {
		"en-US": "Tuesday",
		"es-ES": "Martes",
		"fa-IR": "سه شنبه",
		"it-IT": "Martedi'",
		"ru-RU": "Вторник",
	},
	COMMAND_TEXT_WEDNESDAY: {
		"en-US": "Wednesday",
		"es-ES": "Miercoles",
		"fa-IR": "چهارشنبه",
		"it-IT": "Mercoledi'",
		"ru-RU": "Среда",
	},
	COMMAND_TEXT_THURSDAY: {
		"en-US": "Thursday",
		"es-ES": "Jueves",
		"fa-IR": "پنج شنبه",
		"it-IT": "Giovedi'",
		"ru-RU": "Четверг",
	},
	COMMAND_TEXT_FRIDAY: {
		"en-US": "Friday",
		"es-ES": "Viernes",
		"fa-IR": "جمعه",
		"it-IT": "Venerdi'",
		"ru-RU": "Пятница",
	},
	COMMAND_TEXT_SATURDAY: {
		"en-US": "Saturday",
		
		"fa-IR": "شنبه",
		"it-IT": "Sabato",
		"ru-RU": "Суббота",
	},
	COMMAND_TEXT_SUNDAY: {
		"en-US": "Sunday",
		"es-ES": "Domingo",
		"fa-IR": "یکشنبه",
		"it-IT": "Domenica",
		"ru-RU": "Воскресенье",
	},
	COMMAND_TEXT_DO_NOT_SEND_RECEIPT: {
		"en-US": "Do not send the receipt",
		"es-ES": "No enviar el recibo",
		"fa-IR": "رسید را ارسال نکنید",
		"it-IT": "Non inviare la ricevuta",
		"ru-RU": "Не отправлять квитанцию",
	},
	MESSAGE_TEXT_RECEIPT_WILL_NOT_BE_SENT: {
		"en-US": "You've decided not to send the receipt.",
		"es-ES": "Ha decidido no enviar el recibo",
		"fa-IR": "شما تصمیم گرفتید که رسید را ارسال نکنید.",
		"it-IT": "Hai scelto di non inviare la ricevuta",
		"ru-RU": "Вы решили не отправлять квитанцию.",
	},
	COMMAND_TEXT_I_HAVE_CHANGED_MY_MIND: {
		"en-US": "I've changed my mind",
		"es-ES": "He cambiado de opinion",
		"fa-IR": "نظرم را عوض کردم.",
		"it-IT": "Ho cambiato idea",
		"ru-RU": "Я передумал(а)",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TELEGRAM: {
		"en-US": "Send by Telegram",
		"es-ES": "Enviar a través de Telegram",
		"fa-IR": "با تلگرام ارسال شود",
		"it-IT": "Invia tramite Telegram",
		"ru-RU": "Отправить через Telelgram",
	},
	COMMAND_TEXT_COUNTERPARTY_HAS_NO_TELEGRAM: {
		"en-US": "Send by FB, WhatsApp, Viber, etc.",
		"es-ES": "Enviar a través de FB, WhatsApp, Viber, etc.",
		"fa-IR": "با فیسبوک، واتس آپ، وایبر و ... ارسال شود.",
		"it-IT": "Invia con FB, WhatsCrap, Viber, etc.",
		"ru-RU": "Отправить через Viber, VK, FB, ...",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_SMS: {
		"en-US": "Send by SMS",
		"es-ES": "Enviar a través de SMS",
		"fa-IR": "با پیام کوتاه ارسال شود",
		"it-IT": "Invia tramite SMS",
		"ru-RU": "Отправить через SMS",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_VK: {
		"en-US": "Send throw VK.com",
		
		"it-IT": "Invia tramite VK.com (Facebook russo)",
		"fa-IR": "ارسال شود VK.com از طریق ",
		"ru-RU": "Отправить через ВКонтакте",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_OK: {
		"en-US": "Send throw OK",
		"es-ES": "Enviar a través de OK",
		"it-IT": "Invia tramite OK",
		"fa-IR": "ارسال شود OK از طریق ",
		"ru-RU": "Отправить через Одноклассники",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_FB: {
		"en-US": "Send throw Facebook",
		"es-ES": "Enviar a través de Facebook",
		"it-IT": "Invia tramite Facebook",
		"fa-IR": "از طریق فیسبوک ارسال شود.",
		"ru-RU": "Отправить через Facebook",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TWT: {
		"en-US": "Send throw Twiter",
		"es-ES": "Enviar a través de Twitter",
		"fa-IR": "از طریق توئیتر ارسال شود.",
		"it-IT": "Invia tramite Twitter",
		"ru-RU": "Отправить через Twitter",
	},
	COMMAND_TEXT_CANCEL_SENDING_RECEIPT_BY_TELEGRAM: {
		"en-US": "Cancel sending receipt by Telegram",
		"es-ES": "Cancelar el envío a través de Telegram",
		"fa-IR": "ارسال رسید با تلگرام کنسل شود",
		"it-IT": "Annulla l'invio tramite Telegram",
		"ru-RU": "Отменить отправку через Telegram",
	},
	COMMAND_TEXT_MAIN_MENU_TITLE: {
		"en-US": "Main /menu",
		"es-ES": "Inicio /menu",
		"fa-IR": "/منو ی اصلی ",
		"it-IT": "Menu' /menu",
		"ru-RU": "Главное /меню",
	},
	MESSAGE_TEXT_NOTHING_TO_CANCEL: {
		"en-US": "Nothing to cancel.",
		"es-ES": "No hay nada que anular.",
		"fa-IR": "چیزی برای کنسل شدن وجود ندارد",
		"it-IT": "Nulla da annullare.",
		"ru-RU": "Нечего отменять.",
	},
	MESSAGE_TEXT_TRANSFER_CREATION_CANCELED: {
		"en-US": "Creation of debt record has been canceled.",
		"es-ES": "Creación recordatorio ha cancelado.",
		"fa-IR": "ایجاد سابقه بدهی کنسل شد.",
		"it-IT": "Creazione record annullata",
		"ru-RU": "Создание записи о долге отменено.",
	},
	COMMAND_TEXT_SHOW_ALL_CONTACTS: {
		"en-US": "Show all...",
		"es-ES": "Mostrar todo...",
		"fa-IR": "نمایش تمام موارد ...",
		"it-IT": "Mostra tutto...",
		"ru-RU": "Показать все...",
	},
	COMMAND_TEXT_CONTACTS: {
		"en-US": "Contacts",
		"es-ES": "Contactos",
		"fa-IR": "لیست تماس",
		"it-IT": "Сontatti",
		"ru-RU": "Контакты",
	},
	COMMAND_TEXT_REFRESH: {
		"en-US": "Refresh",
		"es-ES": "Recargar",
		"fa-IR": "تازه کردن",
		"it-IT": "Ricaricare",
		"ru-RU": "Обновить",
	},
	COMMAND_TEXT_ADD_YOUR_OWN_OPTION: {
		"en-US": "Something else",
		"es-ES": "Otra cosa",
		"fa-IR": "چیزی دیگر",
		"it-IT": "Qualcos'altro",
		"ru-RU": "Что-то другое",
	},
	MESSAGE_TEXT_REMINDER_ASK_IF_RETURNED: {
		"en-US": "Have this debt been returned?",
		"es-ES": "¿Se ha devuelto esta deuda?",
		"fa-IR": "آیا این بدهی بازپرداخت شده است؟",
		"it-IT": "Questo debito e' stato saldato?",
		"ru-RU": "Был ли этот долг возвращён?",
	},
	MESSAGE_TEXT_ASK_WHEN_TO_REMIND_AGAIN: {
		"en-US": "When should we remind you about this debt again?",
		"es-ES": "¿Cuándo recordarte de esta deuda otra vez?",
		"it-IT": "Quando devo ricordarti di questo debito?",
		"fa-IR": "چه زمانی لازم است مجدداً در مورد این بدهی به شما یادآوری نماییم؟",
		"ru-RU": "Когда вам напомнить об этом долге ещё раз?",
	},
	MESSAGE_TEXT_REPLIED_DEBT_RETURNED_FULLY: {
		"en-US": "You've replied back that debt has been returned fully.",
		"es-ES": "Has confirmado que la deuda se volvió totalemente",
		"it-IT": "Hai confermato che il debito e' stato saldato.",
		"fa-IR": "شما پاسخ داده اید که بدهی به صورت کامل بازپرداخت شده است.",
		"ru-RU": "Вы ответили что долг возвращён полностью.",
	},
	MESSAGE_TEXT_DEBT_IS_RETURNED: {
		"en-US": "The debt has been returned fully.",
		"es-ES": "La deuda ha devuelto totalemente",
		"fa-IR": "بدهی به صورت کامل بازپرداخت شده است",
		"it-IT": "Il debito e' stato saldato.",
		"ru-RU": "Долг возвращён полностью.",
	},
	MESSAGE_TEXT_DETAILS_ARE_HERE: {
		"en-US": "Details here: %v",
		"es-ES": "Detalles aquí",
		"fa-IR": "جزئیات در اینجا: %v",
		"it-IT": "Dettagli qui: %v",
		"ru-RU": "Подробности тут: %v",
	},
	MESSAGE_TEXT_REMINDER: {
		"en-US": "Reminder",
		"es-ES": "Recordatorio",
		"fa-IR": "یادآور",
		"it-IT": "Promemoria",
		"ru-RU": "Напоминание",
	},
	MESSAGE_TEXT_REMINDER_SET: {
		"en-US": "Reminder set for: %v",
		"es-ES": "Recordatorio establecito para: %v",
		"fa-IR": "یادآور تنظیم شده است برای: %v",
		"it-IT": "Imposta promemoria per: %v",
		"ru-RU": "Напоминание установлено на: %v",
	},
	MESSAGE_TEXT_REMINDER_DISABLED: {
		"en-US": "You've disabled reminders for this debt.",
		"es-ES": "Recordatorio para esta deuda se ha desconectado.",
		"fa-IR": "شما یادآور را برای این بدهی غیرفعال نموده اید.",
		"it-IT": "Hai disabilitato il promemoria per questo debito.",
		"ru-RU": "Напоминания об этом долге отключены.",
	},
	MESSAGE_TEXT_REMINDER_ALREADY_RESCHEDULED: {
		"en-US": "You've already rescheduled this reminder.",
		"es-ES": "Recordatorio para esta deuda se ha reprogramado ya.",
		"it-IT": "Hai gia' impostato questo promemoria",
		"fa-IR": "شما قبلا به صورت مجدد این یادآور را زمانبندی نموده اید.",
		"ru-RU": "Напоминание об этом долге уже перенесено.",
	},
	COMMAND_TEXT_REMINDER_RETURNED_IN_FULL: {
		"en-US": "Yes, returne in full",
		"es-ES": "Sí, devuelto totalmente",
		"fa-IR": "بله، بازپرداخت به صورت کامل",
		"it-IT": "Si, completamento saldato",
		"ru-RU": "Да, возвращено полностью",
	},
	COMMAND_TEXT_REMINDER_RETURNED_PARTIALLY: {
		"en-US": "Returned partially",
		"es-ES": "Devuelto patcialmente",
		"fa-IR": "تا اندازه ای بازپرداخت شده است",
		"it-IT": "Parzialmente saldato",
		"ru-RU": "Возврашено частично",
	},
	COMMAND_TEXT_REMINDER_NOT_RETURNED: {
		"en-US": "Not returned",
		"es-ES": "No devuelto",
		"fa-IR": "بازپرداخت نشده است",
		"it-IT": "Debito non saldato",
		"ru-RU": "Не возвращено",
	},
	MESSAGE_TEXT_YOU_REPLIED: {
		"en-US": "You've replied: %v",
		"es-ES": "Has devuelto: %v",
		"fa-IR": "شما پاسخ داده اید: %v",
		"it-IT": "Hai risposto: %v",
		"ru-RU": "Вы ответили: %v",
	},
	"book": {
		"en-US": "book",
		"es-ES": "libro",
		"fa-IR": "کتاب",
		"it-IT": "libro",
		"ru-RU": "книгу",
	},
	bots.MESSAGE_TEXT_I_DID_NOT_UNDERSTAND_THE_COMMAND: {
		"en-US": "\xF0\x9F\x98\xB3 Sorry, I did not understand your order. May be I'm a little bit dumb...\n\nYou can return to main /menu",
		"es-ES": "\xF0\x9F\x98\xB3 Disculpa, no he entendido tu mando. Tal vez soy un poco romo...\n\nPPuedes volver al Inicio /menu",
		"fa-IR": "\xF0\x9F\x98\xB3 ببخشید، من دستور شما را نفهمیدم. احتمالا کمی کند ذهن هستم...\n\nشما میتوانید به /منو ی اصلی بازگردید",
		"it-IT": "\xF0\x9F\x98\xB3 Scusami ma non ho capito cosa vuoi. Sono ancora un po' sciocco...\n\nPuoi ritornare al Menu con /menu",
		"ru-RU": "\xF0\x9F\x98\xB3 Извините, я не понял вашу команду. Возможно я немного туповат...\n\nЧтобы начать сначала нажмите /menu",
	},
	"COMMAND_TEXT_LANGUAGE": {
		"en-US": "App /language",
		"es-ES": "App /Idioma",
		"fa-IR": "App /زبان",
		"it-IT": "/Lingua",
		"ru-RU": "/Язык приложения",
	},
	"/start": {
		"en-US": "/start",
		"es-ES": "/comienzo",
		"fa-IR": "/شروع",
		"it-IT": "/start",
		"ru-RU": "/старт",
	},
	COMMAND_TEXT_DUE_RETURNS: {
		"en-US": "Due returns",
		"es-ES": "Devoluciones",
		"fa-IR": "بازپرداخت بدهی",
		"it-IT": "Debiti",
		"ru-RU": "Предстоящие платежи",
	},
	MESSAGE_TEXT_OVERDUE_RETURNS_HEADER: {
		"en-US": "<b>Overdue debts:</b>",
		"es-ES": "Deudas pendientes",
		"fa-IR": "<b>بدهی های معوق:</b>",
		"it-IT": "<b>Debiti in ritardo:</b>",
		"ru-RU": "<b>Просроченные долги:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_HEADER: {
		"en-US": "<b>Closest debts to return:</b>",
		"es-ES": "Deudas más cercanos a regresar:</b>"",
		"fa-IR": "<b>نزدیک ترین بدهی برای بازپرداخت:</b>",
		"it-IT": "<b>Debiti in scadenza:</b>",
		"ru-RU": "<b>Ближайшие долги к возрату:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_USER: {
		"en-US": "%v expects %v from you in %v",
		"es-ES": "%v espera %v que devuelvas en %v",
		"it-IT": "%v aspetta %v da te entro il %v",
		"fa-IR": "%v انتظار دارد %v از شما در %v",
		"ru-RU": "%v ожидает от вас возврата %v через %v",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_COUNTERPARTY: {
		"en-US": "You expect %v to retun %v to you in %v",
		"es-ES": "%v esperas de%v que devuelva en %v",
		"fa-IR": "شما انتظار دارید %v بازگرداند %v به شما در %v",
		"it-IT": "Stai aspettando %v che ti dia %v entro il %v",
		"ru-RU": "Вы ожидаете от %v возврата %v через %v",
	},
	MESSAGE_TEXT_DUE_RETURNS_EMPTY: {
		"en-US": "You have no debts with set due date.",
		"es-ES": "No tienes deudas con la fecha señalada para volver. ",
		"fa-IR": "شما بدهی ای با ثبت سررسید ندارید.",
		"it-IT": "Non hai debiti con una data di scadenza.",
		"ru-RU": "У вас нет долгов с назначеным сроком к возврату.",
	},
	COMMAND_TEXT_GAVE: {
		"en-US": "/Gave",
		"es-ES": "/Crédito",
		"fa-IR": "/قرض_دادن",
		"it-IT": "/Credito",
		"ru-RU": "/Дал",
	},
	COMMAND_TEXT_GOT: {
		"en-US": "/Got",
		"es-ES": "/Débito",
		"fa-IR": "/قرض_گرفتن",
		"it-IT": "/Debito",
		"ru-RU": "/Взял",
	},
	COMMAND_TEXT_RETURN: {
		"en-US": "/Returned",
		
		"fa-IR": "/بازگشت",
		"it-IT": "/Rientra",
		"ru-RU": "/Вернул",
	},
	COMMAND_TEXT_BALANCE: {
		"en-US": "/Balance",
		"es-ES": "/Balance",
		"fa-IR": "/تراز",
		"it-IT": "/Bilancio",
		"ru-RU": "/Баланс",
	},
	COMMAND_TEXT_SETTING: {
		"en-US": "/Settings",
		"es-ES": "/Ajustes",
		"fa-IR": "/تنظیمات",
		"it-IT": "/Settaggi",
		"ru-RU": "/Настройки",
	},
	COMMAND_TEXT_HIGH_FIVE: {
		"en-US": "High five!",
		"es-ES": "Cinco altos!",
		"fa-IR": "بزن قدش!",
		"it-IT": "Batti 5 bro!",
		"ru-RU": "Дать пять!",
	},
	COMMAND_TEXT_CHANGE_LANG: {
		"en-US": "/Language",
		"es-ES": "/Idioma",
		"fa-IR": "/زبان",
		"it-IT": "/Lingua",
		"ru-RU": "/Язык",
	},
	COMMAND_TEXT_HELP: {
		"en-US": "/Help",
		"es-ES": "/Ayuda",
		"fa-IR": "/کمک",
		"it-IT": "/Aiuto",
		"ru-RU": "/Помощь",
	},
	COMMAND_TEXT_HISTORY: {
		"en-US": "/History",
		
		"fa-IR": "/پیشینه",
		"it-IT": "/Cronologia",
		"ru-RU": "/История",
	},
	COMMAND_TEXT_CANCEL: {
		"en-US": "/Cancel",
		"es-ES": "/Cancelar",
		"fa-IR": "/کنسل",
		"it-IT": "/Annulla",
		"ru-RU": "/Отменить",
	},
	BUTTON_TEXT_CANCEL: {
		"en-US": "↩ Cancel",
		"es-ES": "/Cancelar",
		"fa-IR": "↪ کنسل",
		"it-IT": "↩ Annulla",
		"ru-RU": "↩ Отменить",
	},
	BUTTON_TEXT_MAIN_MENU: {
		"en-US": "↩ Main menu",
		"ru-RU": "↩ Главное меню",
		//"fa-IR": "↪ ", // TODO(FA)
		//"it-IT": "↩ ", // TODO(IT)
	},
	COMMAND_TEXT_SETTINGS_PRIMARY_CURRENCY: {
		"en-US": "Primary currency",
		"es-ES": "Moneda principal",
		"fa-IR": "واحد پول اولیه",
		"it-IT": "Valuta principale",
		"ru-RU": "Основная валюта",
	},
	COMMAND_TEXT_NEW_COUNTERPARTY: {
		"en-US": "Add new",
		"es-ES": "Añadir nuevo",
		"fa-IR": "اضافه کردن مورد جدید",
		"it-IT": "Aggiungi nuovo",
		"ru-RU": "Добавить",
	},
	MESSAGE_TEXT_LOGIN_CODE: {
		"en-US": "Your code for signing in to app: <b>%v</b>",
		"es-ES": "Tu código para entrar en app: <b>%v</b>",
		"fa-IR": "کد شما برای ورود به برنامه: <b>%v</b>",
		"it-IT": "Il tuo codice per accedere all'app e': <b>%v</b>",
		"ru-RU": "Ваш код для входа в приложение: <b>%v</b>",
	},
	MESSAGE_TEXT_ASK_NEW_COUNTERPARTY_NAME: {
		"ru-RU": `<b>Имя для нового контакта</b>
		Напишите сами или выберите из своей адресной книги (<i>через иконку "скрепка"</i>).

		<i>Отправьте '.' для отмены</i>`,
		"en-US": `Please enter a name for the new contact:
		You can type manually or choose from your address book (<i>through "clip" icon</i>).

		<i>Send '.' to cancel</i>`,

		"fa-IR": `لطفا برای مخاطب جدید یک نام وارد کنید:
		میتوانید به صورت دستی تایپ نموده یا از لیست مخاطبین خود انتخاب نمایید (<i>throw "clip" icon</i>).

		<i>Send '.' برای کنسل کردن</i>`,

		"it-IT": `Inserisci un nome per il nuovo contatto:
		Puoi digitarlo o sceglierlo dalla tua rubrica (<i>attraverso l'icona "clip"</i>).

		<i>Digita '.' ed invia per annullare</i>`,

	},
	MESSAGE_TEXT_TRANSFER_IS_CREATING: {
		"ru-RU": "Создаю запись...",
		"en-US": "Creating transfer...",
		"fa-IR": "ایجاد انتقال ...",
		"it-IT": "Sto creando la nuova voce...",
                "es-ES": "Creando la nueva nota...",
	},
	COMMAND_TEXT_PLEASE_WAIT: {
		"ru-RU": "Пожалуйста подождите",
                "es-ES": "Espera, por favor",
		"en-US": "Please wait",
		"fa-IR": "لطفا صبر کنید",
		"it-IT": "Aspetta per favore",
	},
	MESSAGE_TEXT_PLEASE_WAIT: {
		"ru-RU": "Пожалуйста подождите...",
                "es-ES": "Espera, por favor...",
		"en-US": "Please wait...",
		"it-IT": "Aspetta per favore...",
		"fa-IR": "لطفا صبر کنید ...",
	},
	MESSAGE_TEXT_SELF_ACKNOWLEDGEMENT: {
		"ru-RU": "Подтверждение ожидается от %v",
                "es-ES": "Confirmación se espera de %v",
		"en-US": "Acknowledgement is expected from %v",
		"it-IT": "Conferma in attesa da %v",
		"fa-IR": "انتظار تصدیق می رود از %v",
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_YOU: {
		"ru-RU": "Вы подтвердили эту транзакцию.",
                "es-ES": "Has confirmado esta transacción",
		"en-US": "You've accepted this transaction.",
		"fa-IR": ".شما این تراکنش را قبول کردید ",
		"it-IT": "Hai accettato il debito/credito.",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_YOU: {
		"ru-RU": `Вы НЕ согласны с этой транзакцией.

Сама транзакция НЕ будет отменена, но создатель будет оповещён.`,
                "es-ES": "No estas de acuerdo con la transacción.
La transacción No será cancelada, pero 
		"en-US": `You do not agree with this transaction.

The transaction will not be deleted but the creator will be notified.`,
		"fa-IR": ".شما این تراکنش را رد کردید", //TODO(FA)
		"it-IT": `Hai rifiutato il debito/credito.

The transaction will NOT be deleted but the creator will be notified.
		`, //TODO(IT)
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_COUNTERPARTY: {
		"ru-RU": "%v подтвердил(a) вашу транзакцию:",
		"en-US": "%v accepted your transaction:",
		"fa-IR": ": تراکنش شمارا تایید کرد %v ",
		"it-IT": "%v ha accettato il tuo credito/debito:",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_COUNTERPARTY: {
		"ru-RU": "%v <b>НЕ</b> подтвердил(a) вашу транзакцию. Транзакция не отменена, но возможно вам стоит это обсудить.",
		"en-US": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(FA)
		"fa-IR": "تراکنش شما را رد کرد  %v declined your transaction.",
		"it-IT": "%v ha rifiutato il tuo credito/debito.  The transaction is not canceled but you may want to discuss it.", //TODO(IT)
	},
	COMMAND_TEXT_SUBSCRIBE_TO_APP: {
		"ru-RU": "Хочу приложение!",
		"en-US": "I want the app!",
		"fa-IR": "!من برنامه را می خواهم",
		"it-IT": "Voglio l'aplicazione!",
	},
	COMMAND_TEXT_I_AM_FINE_WITH_BOT: {
		"ru-RU": "Меня вполне устраивает бот!",
		"en-US": "I'm fine with just the bot!",
		"fa-IR": "! ربات به تنهایی برای من کافی است",
		"it-IT": "Mi accontento del bot per ora!",
	},
	MESSAGE_TEXT_SUBSCRIBED_TO_APP: {
		"ru-RU": "Мы сообщим вам когда приложение будет доступно для загруки.",
		"en-US": "We'll let you know once the app is available for download.",
		"fa-IR": ".وقتی برنامه برای دانلود دردسترس بود به شما اطلاع می دهیم",
		"it-IT": "Ti faremo sapere non appena l'applicazione sara' disponibile al download.",
	},
	MESSAGE_TEXT_NOT_INTERESTED_IN_APP: {
		"ru-RU": "Чтож, мы рады что вас устраивает наш бот и нет необходимости загружать приложение.",
		"en-US": "Well, we are happy our bot is good enough and there is no need to download an app.",
		"fa-IR": ".خب، ما خوشحال هستیم که ربات برای شما کافی است و نیازی به دانلود برنامه نیست",
		"it-IT": "Bene, siamo contenti che il nostro bot sia di tuo gradimento e non hai bisogno di scaricare l'applicazione.",
	},
	MESSAGE_TEXT_YOUR_AD_COULD_BE_HERE: {
		"ru-RU": "Здесь можно <a href>разместить рекламу</a>",
		"en-US": "You can <a href>advertise here</a>",
		"fa-IR": "شما میتوانید <a href>در اینجا تبلیغ کنید</a>",
		"it-IT": "Puoi <a href>pubblicizzare qui</a>",
	},
	MESSAGE_TEXT_YOUR_ABOUT_ADS: {
		"ru-RU": `🤖: Я конечно обоятельный робот, но пользоваться специализированным приложением бывает удобнее. Оно ещё не готово для общего доступа, но уже сейчас можно посмотреть как будет выглядеть: <a href= "https://debtstracker.io/ru/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ru/</a>

		Хотите получить оповещение когда оно выйдет?`,

		"en-US": `🤖: I'm a good robot, for sure. But sometimes it is more convinient to use a nice specialized app. It's not ready for public use yet but you can check how it is going to looks: <a href = "https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	Do you want to get an invite when it gets released?`,

		"it-IT": `🤖: Di sicuro son un bravo bot, ma alcune volte e' piu' conveniente usare un'applicazione specializzata. Non e' ancora pronta per la pubblicazione ma puoi controllare l'avanzamento a questo indirizzo: <a href = "https://debtstracker.io/it/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/it/</a>

	Vuoi essere invitato non appena viene rilasciata?`,

		"fa-IR": `🤖: مطمئناً من روبات خوبی هستم. اما بعضی وقت هاساده تر و مناسب تر است که از یک برنامه به خوبی تخصصی شده استفاده شود، این برنامه هنوز برای استفاده عموم آماده نیست ولی می توانید چک کنید که چگونه به نظر خواهد رسید: <a href = "https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	آیا می خواهید وقتی منتشر شد دعوتنامه ای دریافت کنید؟`,
	},
	MESSAGE_TEXT_INVALID_FLOAT: {
		"ru-RU": "Извините, но вы можете использовать только числа в качестве суммы/количества (<i>до 2х знаком после запятой</i>).",
		"en-US": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		"fa-IR": "ببخشید، اما شما تنها میتوانید از اعداد بعنوان مقادیر / اندازه ها استفاده کنید (<i>با دو رقم اعشار</i>).",
		"it-IT": "Spiacente, puoi utilizzare solo numeri (<i>con un massimo di 2 numeri dopo il punto</i>).",
	},
	MESSAGE_TEXT_ASK_LENDING_TYPE: {
		"ru-RU": "<b>Что вы дали в долг?</b>",
		"en-US": "<b>What did you lend to someone?</b>",
		"fa-IR": "<b> چه چیزی به کسی قرض داده اید؟</b>",
		"it-IT": "<b>Cos'hai prestato?</b>",
	},
	MESSAGE_TEXT_CHOOSE_CURRENCY: {
		"ru-RU": `Выберите из меню внизу экрана или <a>выберите валюту из списка</a>.

	Если ни один из стандартных вариантов не подходит просто напишите текстом.Например: "<i>яблоко</i>".`,

		"en-US": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		"fa-IR": `لطفا از بین گزینه های زیر انتخاب کنید یا <a>یک واحد پولی از لیست انتخاب کنید</a>.

	اگر گزینه های استاندارد کافی نبودند به سادگی یک متن بفرستید ، برای مثال:. "<i>سیب</i>".`,

		"it-IT": `Scegli dalle opzioni qui sotto o <a>seleziona una valuta dalla lista</a>.

	Se le opzioni standard non bastano semplicemente invia un testo.Per esempio: "<i>mele</i>".`,
	},
	MESSAGE_TEXT_ASK_LENDING_AMOUNT: {
		"ru-RU": "Сколько <b>%v</b> вы дали в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		"fa-IR": "چه مقدار <b>%v</b> قرض داده اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		"it-IT": "Quanti <b>%v</b> hai prestato?\n(<i>Digita '.' ed invia per annullare</i>)",
	},
	MESSAGE_TEXT_ASK_LENDING_COUNTERPARTY: {
		"ru-RU": "Кому вы дали в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		"fa-IR": "چه کسی از شما <b>%v</b> قرض گرفته است؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		"it-IT": "Chi e' in debito di <b>%v</b> con te?\n(<i>Digita '.' ed invia per annullare</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_TYPE: {
		"ru-RU": "Что вы взяли в долг?",
		"en-US": "What did you lend?",
		"fa-IR": "چه چیزی قرض گرفته اید؟",
		"it-IT": "Cosa ti hanno prestato?",
	},
	MESSAGE_TEXT_ASK_BORROWING_AMOUNT: {
		"ru-RU": "Сколько <b>%v</b> вы взяли в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		"fa-IR": "چه مقدار <b>%v</b> قرض گرفته اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		"it-IT": "Quanti <b>%v</b> ti hanno prestato?\n(<i>Digita '.' ed invia per annullare</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_COUNTERPARTY: {
		"ru-RU": "У кого вы взяли в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"en-US": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		"fa-IR": "چه کسی به شما <b>%v</b> قرض داده است؟ \n(<i>ارسال '.' برای کنسل کردن</i>)",
		"it-IT": "Chi ti ha prestato <b>%v</b>?\n(<i>Digita '.' ed invia per annullare</i>)",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT: {
		"ru-RU": "Отправить <a receipt>квитанцию</a> для <a counterparty>%v</a>?",
		"en-US": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		"fa-IR": "آیا لازم است ماارسال کنیم یک <a receipt>رسید</a> به <a counterparty>%v</a>?",
		"it-IT": "Devo inviare una <a receipt>notifica</a> a <a counterparty>%v</a>?",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT_TO_YOURSELF_BY_SMS: {
		"ru-RU": "К сожалению отправка квитанцию себе по СМС в данный момент отключена. Но вы можете отправить её для %v.",
		"en-US": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		"fa-IR": "متاسفم، درحال حاضرارسال یک رسید به خودتان بوسیله پیام کوتاه امکان پذیر نیست. شما میتوانید آنرا ارسال کنید به  %v از طریق.",
		"it-IT": "Spiacente ma inviarsi da soli una notifica tramite SMS non e' al momento disponibile. Pero' puoi inviarla a %v.",
	},
	MESSAGE_TEXT_RECEIPT_IS_SENDING_BY_TELEGRAM: {
		"ru-RU": "Отправляем для %v извещение через Telegram...",
		"en-US": "We are sending receipt to %v by Telegram...",
		"fa-IR": "مادرحال ارسال رسید به %v از طریق تلگرام هستیم...",
		"it-IT": "Sto inviando la notifica a %v tramite Telegram...",
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_FROM_USER: {
		"ru-RU": "{{.Counterparty}} взял(а) в долг {{.Amount}}.",
		"en-US": "{{.Counterparty}} borrowed from you {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} از شما {{.Amount}} قرض گرفته است .",
		"it-IT": "{{.Counterparty}} e' in debito di {{.Amount}} con te.",
		//"it-IT": "{{.Counterparty}} ha preso in prestito da te {{.Amount}}.",

	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_TO_USER: {
		"ru-RU": "{{.Counterparty}} дал(а) вам в долг {{.Amount}}.",
		"en-US": "{{.Counterparty}} lended to you {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} به شما {{.Amount}} قرض داده است .",
		"it-IT": "{{.Counterparty}} ti ha prestato {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_FROM_USER: {
		"ru-RU": "Вы вернули долг - {{.Counterparty}} получил(а) {{.Amount}}.",
		"en-US": "You returned {{.Amount}} to {{.Counterparty}}.",
		"fa-IR": "شما بازگردانده اید {{.Amount}} به {{.Counterparty}}.",
		"it-IT": "Hai ridato {{.Amount}} a {{.Counterparty}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_TO_USER: {
		"ru-RU": "{{.Counterparty}} вернул вам {{.Amount}}.",
		"en-US": "{{.Counterparty}} returned to you {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} به شما بازپرداخت کرده است {{.Amount}}.",
		"it-IT": "{{.Counterparty}} ti ha ridato {{.Amount}}.",
	},
	MESSAGE_TEXT_TRANSFER_ALREADY_FULLY_RETURNED:  {
		"ru-RU": "Этот долг уже полностью возвращён.",
		"en-US": "This debts is already fully returned.",
		//"fa-IR": "", TODO(FA)
		//"it-IT": "", TODO(IT)
	},
	MESSAGE_TEXT_RECEIPT_ALREADY_RETURNED_AMOUNT:  {
		"ru-RU": "Уже возвращено: {{.Amount}}.",
		"en-US": "Already returned: {{.Amount}}.",
		//"fa-IR": "", TODO(FA)
		//"it-IT": "", TODO(IT)
	},
	MESSAGE_TEXT_RECEIPT_OUTSTANDING_AMOUNT:  {
		"ru-RU": "Осталось вернуть: {{.Amount}}.",
		"en-US": "Outstanding: {{.Amount}}.",
		//"fa-IR": "", TODO(FA)
		//"it-IT": "", TODO(IT)
	},
	MESSAGE_TEXT_DUE_ON: {
		"ru-RU": "<b>Вернуть до</b>: %v",
		"en-US": "<b>Return till</b>: %v",
		"fa-IR": "<b>بازگردانده شود تا</b>: %v",
		"it-IT": "<b>Scadenza</b>: %v",
	},
	MESSAGE_TEXT_NOTE: {
		"ru-RU": "Заметка",
		"en-US": "Note",
		"fa-IR": "نکته",
		"it-IT": "Nota",
	},
	MESSAGE_TEXT_COMMENT: {
		"ru-RU": "Комментарий",
		"en-US": "Comment",
		"fa-IR": "شرح",
		"it-IT": "Commenti",
	},
	MESSAGE_TEXT_LOGIN_TO_WEB_APP: {
		"ru-RU": `Перейдите по <a>ссылке</a> чтобы запустить web-приложение.`,
		"en-US": `Click to <a>sign in</a> to web-app.`,
		"it-IT": "Fai clic per <a>accedi</a> per app web.",
		"fa-IR": `کلیک کنید تا <a>وارد شوید</a> برنامه وب.`,
	},
	MESSAGE_TEXT_DO_YOU_LIKE_OUR_BOT: {
		"ru-RU": "Вам нравится @{{bot}}?",
		"en-US": "Do you like @{{bot}}?",
		"it-IT": "Divertito con @{{bot}}?",
		"fa-IR": "آیا می پسندید @{{bot}}?",
	},
	COMMAND_TEXT_YES_EXCLAMATION: {
		"ru-RU": "%v Да!",
		"en-US": "%v Yes!",
		"it-IT": "%v Si!",
		"fa-IR": "بله! %v",
	},
	COMMAND_TEXT_YES: {
		"ru-RU": "%v Да",
		"en-US": "%v Yes",
		"it-IT": "%v Si",
		"fa-IR": "بله %v",
	},
	COMMAND_TEXT_NO: {
		"ru-RU": "%v Нет",
		"en-US": "%v No",
		"it-IT": "%v No",
		"fa-IR": "خیر %v",
	},
	COMMAND_TEXT_NOT_TOO_MUCH: {
		"ru-RU": "%v Не очень",
		"en-US": "%v Not too much",
		"it-IT": "%v Non troppo",
		"fa-IR": "نه خیلی زیاد %v",
	},
	COMMAND_TEXT_FEEDBACK: {
		"ru-RU": "/Отзыв",
		"en-US": "/Feedback",
		"it-IT": "/Risposta",
		"fa-IR": "/بازخورد",
	},
	COMMAND_TEXT_WRITE_FEEDBACK: {
		"ru-RU": "%v Написать отзыв",
		"en-US": "%v Write feedback",
		"it-IT": "%v Scrivi commenti",
		"fa-IR": "ارسال بازخورد %v",
	},
	MESSAGE_TEXT_THANKS: {
		"ru-RU": "🙏 Спасибо!",
		"en-US": "🙏 Thanks!",
		"it-IT": "🙏 Grazie!",
		"fa-IR": "🙏 تشکر!",
	},
	MESSGE_TEXT_DEBT_ERROR_FIXED_START_OVER: {
		"en-US": "🙏 Sorry, there was an error. It has been fixed but please re-enter your data for this debt.",
		"ru-RU": "🙏 Извините, у нас была ошибка. Она была исправлено, но потребуется внести данные для этого долга заново.",
	},
	MESSAGE_TEXT_PLEASE_SEND_TEXT: {
		"ru-RU": "Пожалуйста отправьте текст.",
		"en-US": "Please send text.",
		"it-IT": "Si prega di inviare il testo.",
		"fa-IR": "لطفاً متن ارسال کنید.",
	},
	MESSAGE_TEXT_CAN_YOU_RATE_AT_STOREBOT: {
		"ru-RU": `🤖 Можете поставить ему высокую оценку и хороший отзыв в каталоге ботов Store Bot?

	‎Это займет меньше минуты вашего времени! 😇`,
		"en-US": `🤖 Can you rate it high and write a good review in bots catalog Store Bot?

	‎It will take less than a minute of your time! 😇`,
		"it-IT": `🤖 Puoi votarlo in alto e scrivere una buona revisione nel catalogo bot Bot Store?

	Ci vorrà meno di un minuto del tuo tempo! 😇`,
		"fa-IR": `🤖  آیا می توانید در کاتالوگ روباتها در استور بوت امتیاز بالایی داده و اظهار نظر خوبی در مورد این روبات ثبت کنید؟  

این کار کمتر از یک دقیقه از وقت شما را می گیرد! 😇`,
	},
	MESSAGE_TEXT_ASK_TO_WRITE_FEEDBACK_WITHIN_MESSENGER: {
		"ru-RU": "Поделитесь вашими мыслями (на русском или английском) о том, что нужно сделать, чтобы бот стал лучше:",
		"en-US": "‎Share your thoughts (in English or Russian) about what could be done to make the bot better:",
		"it-IT": "Condividi i tuoi pensieri (in Inglese o Russo) su come sarebbe migliore secondo te il bot:",
		"fa-IR": "نظرات خود را (به انگلیسی و روسی ) در مورد اینکه چه کاری می توان انجام داد تا این ربات بهتر شود، با ما به اشتراک بگذارید:",
	},
	MESSAGE_TEXT_HOW_TO_RATE_AT_STOREBOT: {
		"ru-RU": `‎<b>Как поставить оценку в три простых шага:</b>

	1. Перейдите по ссылке, чтобы оставить оценку и отзыв:
	‎https://t.me/storebot?start={{bot}}

	‎2. Нажмите на кнопку "⭐️⭐️⭐️⭐️⭐️"

	‎3. Напишите сообщение или нажмите кнопку "Пропустить этот шаг"

	Спасибо вам большое! Благодаря этому о боте узнает больше людей — это служит дополнительной мотивацией для разработчиков! 😎`,

		"en-US": `<b>How to rate in 3 simple steps:</b>

	1. Click on this link to rate and review:
	https://t.me/storebot?start={{bot}}

	‎2. Click on the "⭐️⭐️⭐️⭐️⭐️" button

	‎3. Write your message or press "Skip this step" button

	Thank you very much! As a result of your actions, even more people will learn about the bot.All this will serve as the additional motivation for the developers! 😎`,

		"it-IT": `<b>Come valutare in 3 semplici passaggi:</b>
	‎1. Clicca su questo link per votare e lasciare una recensione:
	‎https://t.me/storebot?start={{bot}}

	‎2. Clicca sul "⭐️⭐️⭐️⭐️⭐️" bottone

	‎3. Scrivi il tuo messaggio o premi "Salta questo step"

	Grazie infinitamente! Come risultato delle tue azioni, altre persone guarderanno il bot.Dando anche un motivo in più per continuare ai developers! 😎`,

		"fa-IR": `<b>چگونگی امتیازدهی در سه گام ساده :</b>

	1. برای امتیازدهی و ثبت نظرات بر روی لینگ زیر کلیک کنید
	https://t.me/storebot?start={{bot}}

	‎2. بر روی دکمه "⭐️⭐️⭐️⭐️⭐️" کلیک کنید

	‎3. پیام خودرا ثبت کنید یا روی دکمه "پرش از این مرحله" کلیک کنید

	بسیار سپاسگزاریم! عمل شما باعث می شود افراد بیشتری در مورد bot.All بیاموزند. این امر انگیزه مضاعفی به توسعه دهندگان این ربات می دهد ! 😎`,
	},
	MESSAGE_TEXT_ASK_FOR_FEEDBAСK: {
		"ru-RU": "Будем признетельны если вы оцените работу нашего приложения. Это займёт всего несколько секунд.",
		"en-US": "We would appreciate if tell us how we doing. It takes just few seconds.",
		"fa-IR": "سپاسگزار خواهیم بود اگر به ما بگویید کارمان چطور بوده است. این تنها چند ثانیه زمان میبرد.",
		"it-IT": "Ci farebbe piacere se lasciassi un voto per il nostro lavoro. Ti bastano solo alcuni secondi.",
	},
	COMMAND_TEXT_GIVE_FEEDBACK: {
		"ru-RU": "Оценить приложение",
		"en-US": "Rate this bot",
		"fa-IR": "به این ربات امتیاز بدهید",
		"it-IT": "Vota questo bot",
	},
	COMMAND_TEXT_OPEN_STOREBOT_FOR_FEEDBACK: {
		"ru-RU": "Оценить на  @Storebot",
		"en-US": "Leave rating at @Storebot",
		"fa-IR": "امتیاز خود را اینجا وارد کنید @Storebot",
		"it-IT": "Lascia un voto a @Storebot",
	},
	MESSAGE_TEXT_ON_REFUSED_TO_RATE: {
		"ru-RU": `ОК, возможно вы смоежете поставить оценку в другой раз.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Так же будем признательны если вы предложите любые улучшения.
	`,
		/*------------------------------------------------------------*/
		"en-US": `OK, maybe you can rate us another time.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you suggest any improvements.
	`,
		/*------------------------------------------------------------*/
		// TODO(IT): Google translated
		"it-IT": `OK, forse ci puoi valutare un'altra volta.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Apprezzeremo anche se suggeriamo qualche miglioramento.
	`,
		/*------------------------------------------------------------*/
		"fa-IR": `بسیار خوب، ممکن است شما بتوانید زمان دیگری به ما امتیاز بدهید.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	همچنین سپاسگزار خواهیم بود اگر شما هرگونه امکان بهبود را با ما در میان بگذارید.
	`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_POSITIVE: {
		"ru-RU": `Спасибо, мы очень старались!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Так же будем признательны если вы <a suggest-idea>предложите улучшения</a>.
	`,
		/*------------------------------------------------------------*/
		"en-US": `Thanks, we worked hard!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you <a suggest-idea>suggest improvements</a>.
	`,
		/*------------------------------------------------------------*/
		"it-IT": `GRAZIE MILLE, abbiamo lavoro duro!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Sarebbe ancora piu' apprezzatto se ci <a suggest-idea>suggerisci qualche miglioramento</a>.
	`,
		/*------------------------------------------------------------*/
		"fa-IR": `ممنونیم، ما سخت کارکرده ایم!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	ما همچنین سپاسگزار خواهیم بود اگر شما<a suggest-idea> هرگونه امکان بهبود را با ما در میان بگذارید </a>را.
	`,
	},
	MESSAGE_TEXT_YOU_CAN_HELP_BY: {
		"ru-RU": `
	Вы нам очень поможете если:

	* Поставите нам 5⭐ в <a storebot>каталоге ботов</a>.

	* Расскажите о боте своим друзьям.
	Например во <a share-vk>ВКонтакте</a>, < a share-fb>Facebook</a> или <a share-twitter>Twitter</a>.

* Поддержите дальнейшую разработку - <a href ="https://goo.gl/Qhh0yL">€2 через PayPal</a>
`,
		/*------------------------------------------------------------*/
		"en-US": `
You can help us a lot if you:

* Give us 5⭐ at <a storebot>directory of bots</a>.

* Tell about the app to your friends.
For example at <a share-fb>Facebook</a> or <a share-twitter>Twitter</a>.

* Support further development - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
`,
		/*------------------------------------------------------------*/
		// TODO(FA)
		"fa-IR": `:شما به ما کمک بسیاری می کنید اگر

* ثبت بازخورد مثبت در <a storebot>دایرکتوری روبات ها</a>.

* به دوستان خود در مورد برنامه اطلاع رسانی کنید.
برای مثال در <a share-fb>فیسبوک</a> یا <a share-twitter>توئیتر</a>.

* از توسعه بیشتر برنامه پشتیبانی کنید - <a href = "https://goo.gl/Qhh0yL">€2 از طریق پی پال</a> (<i>$2.2 حدود </i>)`,

		// TODO(IT)
		"it-IT": `
		Ci aiuteresti moltissimo se:

		  * Lasci un feedback positivo 5⭐ alla <a storebot>pagina del bot</a>.

		  * Raccontare dell'app ai tuoi amici.
		    Per esempio su <a share-fb>Facebook</a> o su <a share-twitter>Twitter</a>.

		  * Supporta ulteriormente lo sviluppo del bot - <a href="https://goo.gl/Qhh0yL">2€ tramite PayPal</a> (<i>circa $2.2</i>)
		`,
	},
	MESSAGE_TEXT_COUNTERPARTY_HAS_EMPTY_BALANCE: {
		"ru-RU": `Нулевой баланс для %v`,
		"en-US": `Balance is empty for %v`,
		"fa-IR": `تراز خالی است برای %v`,
		"it-IT": `Non hai alcun credito o debito con %v`,
	},
	MESSAGE_TEXT_ASK_TO_TRANSLATE: {
		"ru-RU": `Хотите чтобы наш бот разговаривал на другом языке? Вы можете <a>помочь с переводом</a>.`,
		"en-US": `Do you want our bot to talk in other language? You can <a>help with translation</a>.`,

		"fa-IR": `آیا می خواهید ربات ما به زبان دیگری صحبت کند؟ شما می توانید <a>با ترجمه به ما کمک کنید</a>.`,

		"it-IT": `Vuoi che il nostro bot parli altre lingue? Aiuta con la <a>traduzione</a>.`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEUTRAL: {
		"ru-RU": `Чтож, мы очень старались. Ваша оценка будет передана разработчикам.

Может быть вы <a submit-bug>сообщите нам что не работает</a> или подскажите <a suggest-idea>как можно улучшить</a>?`,
		/*------------------------------------------------------------*/
		"en-US": `Well, we worked hard. You feedback will be passed to developers.

Maybe you can <a submit-bug>report your issue</a> or <a suggest-idea>suggest how we can improve</a>?`,
		/*------------------------------------------------------------*/

		"fa-IR": `خب، ما سخت کارکردیم. بازخورد شما به توسعه دهندگان برنامه منعکس می شود.

شما می توانید <a submit-bug>مشکل خود را گزارش دهید</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,

		"it-IT": `Bene, il nostro lavoro non e' andato in vano. Il tuo feedback sara' inoltrato agli sviluppatori.

Per caso vuoi anche <a submit-bug>segnalare un problema</a> oppure <a suggest-idea>suggerire come possiamo migliorare</a>?`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEGATIVE: {
		"ru-RU": `Нам очень стыдно. Может быть вы <a submit-bug>подскажите что не так</a> или <a suggest-idea>предложите усовершенствования</a>?`,
		/*------------------------------------------------------------*/
		"en-US": `We are very sorry. Maybe you can <a submit-bug>let us know what is wrong</a> or <a suggest-idea>suggest how we can improve</a>?`,
		/*------------------------------------------------------------*/

		"fa-IR": `ما بسیار متاسفیم. شما می توانید <a submit-bug>به ما بگویید مشکلتان چیست</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,

		"it-IT": `Ci dispiace molto. Potresti farci sapere<a submit-bug>cosa non ti e' piaciuto</a> oppure <a suggest-idea>suggerirci come possiamo migliorare</a>?`,
	},
	COMMAND_TEXT_ASK_FOR_FEEDBACK: {
		"ru-RU": "Оцените наше приложение?",
		"en-US": "Please rate our app",
		"fa-IR": "لطفاً به برنامه ما امتیاز دهید",
		"it-IT": "Per favore vota il nostro bot",
	},
	COMMAND_TEXT_FEEDBACK_POSITIVE: {
		"ru-RU": "Да, отличное приложение!",
		"en-US": "Yes, it's a great app!",
		"fa-IR": "بله، این برنامه عالی است",
		"it-IT": "Si, e' un app fantastica!",
	},
	COMMAND_TEXT_FEEDBACK_NEUTRAL: {
		"ru-RU": "Неплохо, но можно лучше.",
		"en-US": "Not bad but can be better.",
		"fa-IR": "بد نیست ولی می تواند بهتر باشد.",
		"it-IT": "Non male ma potrebbe esser migliore.",
	},
	COMMAND_TEXT_FEEDBACK_NEGATIVE: {
		"ru-RU": "Не нравится",
		"en-US": "Don't like it",
		"fa-IR": "از این برنامه را نمی پسندم",
		"it-IT": "Non mi piace",
	},
	COMMAND_TEXT_FEEDBACK_NOT_READY: {
		"ru-RU": "Пока не понятно",
		"en-US": "Not decided yet",
		"fa-IR": "هنوز تصمیم نگرفته ام.",
		"it-IT": "Sono indeciso",
	},
	MESSAGE_TEXT_SETTINGS: {
		"ru-RU": "Что будем настраивать?",
		"en-US": "What do you want to adjust?",
		"fa-IR": "می خواهید چه چیزی را تنظیم کنید؟",
		"it-IT": "Cosa vuoi modificare?",
	},
	MESSAGE_TEXT_NOT_IMPLEMENTED_YET: {
		"ru-RU": "Извините, данный функционал ещё не реализован.",
		"en-US": "Sorry, this functionality is not implemented yet.",
		"fa-IR": "متاسفم، این عملکرد هنوز پیاده سازی نشده است.",
		"it-IT": "Spiacenti ma questa funzionalita' non e' ancora attiva.",
	},
	MESSAGE_TEXT_ASK_INVITE_CHANNEL: {
		"ru-RU": "Как вы хотите получить код приглашения?",
		"en-US": "How do you want to get an invite?",
		"fa-IR": "می خواهید چگونه دعوت شوید؟",
		"it-IT": "Come vuoi ricevere l'invito?",
	},
	MESSAGE_TEXT_PLEASE_ENTER_INVITE_CODE: {
		"ru-RU": "Пожалуйста введите код приглашения:",
		"en-US": "Please enter an invite code:",
		"fa-IR": "لطفاً یک کد دعوت وارد کنید:",
		"it-IT": "Inserisci un codice invito:",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_RECEIVED: {
		"ru-RU": "Мы отправили письмо на %v.\n\nПожалуйста откройте его и кликните на ссылку для подтверждения адреса.",
		"en-US": "We've sent a message to %v.\n\nPlease open the email and click a link to confirm your email address.",
		"fa-IR": "ما یک پیام ارسال کردیم به %v.\n\nلطفاً ایمیل خود را باز کرده و روی لینک کلیک کنید تا آدرس ایمیل شما تایید شود.",
		"it-IT": "Abbiamo inviato un messaggio a %v.\n\nPer favore apri l'email e clicca sul link per confermare il tuo indirizzo email",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_SENT_TELEGRAM: {
		"ru-RU": "После того как откроется Telegram нажмите кнопку <b>Start</b>.",
		"en-US": "Once Telegram pop ups click the <b>Start</b> button.",
		"fa-IR": "وقتی تلگرام اجرا شد برروی دکمه  <b>شروع</b> کلیک کنید.",
		"it-IT": "Una volta aperto il bot su telegram clicca su <b>Start</b>.",
	},
	MESSAGE_TEXT_USER_CONTACT_FOR_INVITE_RECEIVED: {
		"ru-RU": "Спасибо, вы записаны в очередь на получение приглашения.\n\nТекущее время ожидания 2-3 дня.\n\nВы можете получить приглашение сегодня если расскажите о нашем боте на Facebook.",
		"en-US": "Thanks, you've been queued for an invite.\n\nCurrent awaiting time is 2-3 days.\n\nYou can get an invite code today by sharing a link to the bot on Facebook.",
		"fa-IR": "سپاسگزاریم، شما در نوبت دعوت قرار گرفتید\n\nزمان انتظار شما در حال حاضر 2-3 روز می باشد.\n\n شما می توانید با به اشتراک گذاری لینک روبات در فیسبوک امروز یک کد دعوت دریافت کنید. ",
		"it-IT": "Grazie, ora sei in coda per un codice invito.\n\nTempo di attesa medio 2-3 giorni.\n\nPuoi ottenere un codice invito subito condividendo il link al bot su Facebook.",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_EMAIL: {
		"ru-RU": "Пожалуйста напишите ваш email адрес:",
		"en-US": "Please provide your email address",
		"it-IT": "Inserisci il tuo indirizzo email:",
		"fa-IR": "لطفاً آدرس ایمیل خود را وارد کنید.",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_PHONE_NUMBER: {
		"ru-RU": "Пожалуйста напишите номер вашего телефона:",
		"en-US": "Please provide your phone number",
		"it-IT": "Inserisci il tuo numero di telefono:",
		"fa-IR": "لطفاً شماره تلفن خود را وارد نمایید.",
	},
	MESSAGE_TEXT_WRONG_INVITE_CODE: {
		"ru-RU": "Неправильный код приглашения: %v",
		"en-US": "Wrong invite code: %v",
		"fa-IR": "کد دعوت اشتباه است %v",
		"it-IT": "Codice invito: %v errato",
	},
	MESSAGE_TEXT_WRONG_EMAIL: {
		"ru-RU": "Неправильный email адрес.",
		"en-US": "Wrong email address.",
		"fa-IR": "آدرس ایمیل اشتباه است.",
		"it-IT": "L'email inserita e' sbagliata.",
	},
	MESSAGE_TEXT_WRONG_PHONE_NUMBER: {
		"ru-RU": "Неправильный номер телефона.",
		"en-US": "Wrong phone number.",
		"fa-IR": "شماره تلفن اشتباه است",
		"it-IT": "Il numero inserito e' sbagliato.",
	},
	MESSAGE_TEXT_OK_PLEASE_TRY_AGAIN: {
		"ru-RU": "Хорошо, попробуйте ещё раз.",
		"en-US": "Ok, please try again.",
		"fa-IR": "بسیار خوب، لطفا مجدداً سعی کنید.",
		"it-IT": "Ok, prova di nuovo.",
	},

	COMMAND_TEXT_MISTYPE_WILL_TRY_AGAIN: {
		"ru-RU": "Я опечатался, попробую ещё раз.",
		"en-US": "I've mistyped, will try again.",
		"fa-IR": "اشتباه تایپ کردم، دوباره امتحان می کنم.",
		"it-IT": "Ho sbagliato, riprovo.",
	},
	COMMAND_TEXT_TELL_ME_MORE_ABOUT_INVITES: {
		"ru-RU": "Расскажите ка мне об этих кодах",
		"en-US": "Tell me more about the codes",
		"fa-IR": "در مورد کدها بیشتر برای من توضیح دهید.",
		"it-IT": "Ulteriori informazioni riguardo il codice invito.",
	},
	COMMAND_TEXT_INVITE_ME_BY_EMAIL: {
		"en-US": "Send me invite code by email",
		"ru-RU": "Хочу код приглашения на email",
		"fa-IR": "کد دعوت را برای من از طریق ایمیل ارسال کنید.",
		"it-IT": "Inviami il codice invito tramite email",
	},
	COMMAND_TEXT_INVITE_ME_BY_SMS: {
		"en-US": "Send me invite code by SMS",
		"ru-RU": "Хочу код приглашения по SMS",
		"fa-IR": "کد دعوت را برای من از طریق پیام کوتاه ارسال کنید.",
		"it-IT": "Inviami il codice invito tramite SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_EMAIL: {
		"en-US": "Send me new invite code by email",
		"ru-RU": "Новый код приглашения на email",
		"fa-IR": "یک کد دعوت جدیداز طریق ایمیل برای من  ارسال کنید.",
		"it-IT": "Inviami il nuovo codice invito tramite email",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_SMS: {
		"en-US": "Send me new invite code by SMS",
		"ru-RU": "Новый код приглашения по SMS",
		"fa-IR": "یک کد دعوت جدید از طریق پیام کوتاه برای من ارسال کنید.",
		"it-IT": "Inviami il nuovo codice invito tramite SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_TELEGRAM: {
		"en-US": "Send me new invite by Telegram",
		"ru-RU": "Получить приграшение в Telegram",
		"fa-IR": "یک کد دعوت جدید از طریق تلگرام برای من ارسال کنید.",
		"it-IT": "Inviami il nuovo codice invito tramite Telegram",
	},
	MESSAGE_TEXT_UNKNOWN_LANGUAGE: {
		"en-US": "Unknown language. Please choose 1 from the options:",
		"ru-RU": "Незнакомый язык. Пожалуйста выберете один из предоставленных:",
		"fa-IR": "زبان ناشناخته. لطفاً یکی از گزینه ها را انتخاب کنید:",
		"it-IT": "Lingua socnosciuta. Per favore scegline una tra le opzioni:",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY: {
		"en-US": "Unknown counterparty. Please choose <b>Add new</b> if it's a new contact.",
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите <b>Добавить</b> если это новый контакт.",
		"fa-IR": "مخاطب تراکنش شناسایی نشد. <b>یک مورد جدید اضافه کنید</b> اگر این ایشان یک مخاطب جدید هستند.",
		"it-IT": "Destinatario sconosciuto. Scegli <b>Aggiugni nuovo</b> se e' un nuovo contatto.",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY_ON_RETURN: {
		"en-US": "Unknown counterparty. Please choose from the list.",
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите из списка.",
		"fa-IR": "مخاطب تراکنش شناسایی نشد. لطفاً از فهرست انتخاب کنید.",
		"it-IT": "Destinatario sconosciuto. Scegli dalla lista qui sotto.",
	},
	MESSAGE_TEXT_UNKNOWN_DEBT: {
		"en-US": "Unknown debt. Please choose from the list.",
		"ru-RU": "Неизвестный долг. Пожалуйста выберите из списка.",
		"fa-IR": "بدهی ناشناخته است. لطفا از فهرست انتخاب کنید.",
		"it-IT": "Debito sconosciuto. Scegli dalla lista qui sotto.",
	},
	MESSAGE_TEXT_BILL_CARD_HEADER: {
		"en-US": "<b>Bill/purchase</b>: <code>%v</code>",
		"ru-RU": "<b>Cчёт/покупка</b>: <code>%v</code>",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	MESSAGE_TEXT_BILL_CARD_HEADER_WITH_STATUS: {
		"en-US": "<b>Bill/purchase</b>: <code>%v</code> — <i>%v</i>",
		"ru-RU": "<b>Cчёт/покупка</b>: <code>%v</code> — <i>%v</i>",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_COUNT: {
		"en-US": "Members (%v):",
		"ru-RU": "Участники (%v):",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW: {
		"en-US": "{{.N}}. {{.MemberName}}", // Non need to change for LTR
		//"fa-IR": "", // TODO(FA)
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_OWES: {
		"en-US": "{{.N}}. {{.MemberName}} — owes {{.Owes}}",
		"ru-RU": "{{.N}}. {{.MemberName}} — должен {{.Owes}}",
		//"fa-IR": "", // TODO(FA)
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_PAID: {
		"en-US": "{{.N}}. {{.MemberName}} — paid {{.Paid}}",
		"ru-RU": "{{.N}}. {{.MemberName}} — заплатил {{.Paid}}",
		//"fa-IR": "", // TODO(FA)
	},
	MESSAGE_TEXT_BILL_ASK_WHO_PAID: {
		"en-US": "Please choose who paid for the bill:",
		"ru-RU": "Пожалуйста выберите кто заплатил по счёту:",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	MESSAGE_TEXT_STATUS: {
		"en-US": "Status: %v",
		"ru-RU": "Статус: %v",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	BUTTON_TEXT_ADD_MEMBER: {
		"en-US": "Add participant",
		"ru-RU": "Добавить участника",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	BUTTON_TEXT_FINALIZE_BILL: {
		"en-US": "🔓 Lock the bill",
		"ru-RU": "🔓 Закрыть счёт",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	BUTTON_TEXT_EDIT_BILL: {
		"en-US": "✏️ Edit",
		"ru-RU": "✏️ Изменить",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	BUTTON_TEXT_SPLIT_MODE: {
		"en-US": "➗ Split: %v",
		"ru-RU": "➗ Делить: %v",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	MESSAGE_TEXT_SPLIT_LABEL_WITH_VALUE: {
		"en-US": "Split: %v",
		"ru-RU": "Делить: %v",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	STATUS_DRAFT: {
		"en-US": "draft",
		"ru-RU": "черновик",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	SPLIT_MODE_EQUALLY: {
		"en-US": "Equally",
		"ru-RU": "Поровну",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	SPLIT_MODE_PERCENTAGE: {
		"en-US": "Percentage",
		"ru-RU": "В процентах",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	SPLIT_MODE_EXACT_AMOUNT: {
		"en-US": "Exact amounts",
		"ru-RU": "Точные суммы",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	SPLIT_MODE_SHARES: {
		"en-US": "Shares",
		"ru-RU": "В долях",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	BUTTON_TEXT_JOIN: {
		"ru-RU": "➕ Присоедениться",
		"en-US": "➕ Join",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	BUTTON_TEXT_DUE: {
		"ru-RU": "📅 До: %v",
		"en-US": "📅 Due: %v",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	NOT_SET: {
		"en-US": "not set",
		"ru-RU": "не задано",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	BUTTON_TEXT_MANAGE_MEMBERS: {
		"ru-RU": "👫 Участники",
		"en-US": "👫 Participants",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	BUTTON_TEXT_CHANGE_BILL_PAYER: {
		"ru-RU": "🔀 Сменить плательщика",
		"en-US": "🔀 Change payer",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	COMMAND_TEXT_I_PAID: {
		"ru-RU": "Я заплатил",
		"en-US": "I paid",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	COMMAND_TEXT_I_OWE: {
		"ru-RU": "Я должен",
		"en-US": "I owe",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	COMMAND_TEXT_OWED_TO_ME: {
		"en-US": "Owed to me",
		"ru-RU": "Должны мне",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	MESSAGE_TEXT_BILL_HEADER: {
		"en-US": "Bill: %v",
		"ru-RU": "Cчёт: %v",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	MESSAGE_TEXT_NEW_DEBT_HEADER: {
		"en-US": "Bill: %v",
		"ru-RU": "Cчёт: %v",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	MESSAGE_TEXT_ALREADY_BILL_MEMBER: {
		"en-US": "%v, you are sharing this bill already.",
		"ru-RU": "%v, вы уже входите в состав участников.",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	MESSAGE_TEXT_USER_JOINED_BILL: {
		"en-US": "%v joined to bill sharing.",
		"ru-RU": "%v присоеденился к совместной оплате.",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	MESSAGE_TEXT_YOU_JOINED_BILL: {
		"en-US": "You've joined to bill sharing.",
		"ru-RU": "Вы присоеденились к совместной оплате.",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	ARTICLE_TITLE_SPLIT_BILL: {
		"ru-RU": "Разделить счёт/покупку",
		"en-US": "Split bill/purchase",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	ARTICLE_SUBTITLE_SPLIT_BILL: {
		"ru-RU": "Сумма: %v\nПоделить траты между друзьями и отследить возвраты.",
		"en-US": "Amount: %v\nShare expenses between friends and track paybacks.",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},

	ARTICLE_NEW_DEBT_TITLE: {
		"ru-RU": "Новый долг",
		"en-US": "New debt",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	ARTICLE_NEW_DEBT_SUBTITLE: {
		"ru-RU": "Сумма: %v\nЗапись долга и рассылка оповещений в день возврата.",
		"en-US": "Amount: %v\nRecord debt and get/send notifications on due date.",
		//"fa-IR": "", // TODO(FA)
		//"it-IT": "", // TODO(IT)
	},
	MESSAGE_TEXT_HI: { // This is the same for all languages.
		"ru-RU": `¡Hola! Hi! Привет! سلام!`,
		"en-US": `¡Hola! Hi! Привет! سلام!`,

		"fa-IR": `¡Hola! Hi! Привет! سلام!`,

		"it-IT": `¡Hola! Hi! Привет! سلام!`,
	},
	MESSAGE_TEXT_BACK_TO_MAIN_MENU: {
		"ru-RU": `Можно вернуться назад в главное /меню`,
		"en-US": `You can go back to main /menu`,
		"fa-IR": `شما میتوانید به /منو ی اصلی مراجعه کنید.`,
		"it-IT": `Puoi tornare al menu' principale tramite /menu`,
	},
	MESSAGE_TEXT_YOUR_SELECTED_PREFERRED_LANGUAGE: { // This is the same for all languages.
		"ru-RU": `Выбранный язык программы: %v`,
		"en-US": `Preferred app language: %v`,
		"fa-IR": `زبان برنامه: %v`,
		"it-IT": `Lingua del bot preferita: %v`,
	},
	MESSAGE_TEXT_ONBOARDING_ASK_TO_CHOOSE_LANGUAGE: {
		"ru-RU": `<b>%v</b>, на каком языке вы хотели бы общаться?
(What is your preferred language?)`,
		"en-US": `<b>%v</b>, what is your preferred language?`,
		"fa-IR": `<b>%v</b>, شما چه زبانی را ترجیح می دهید؟
(What is your preferred language?)`,
		"it-IT": `<b>%v</b> qual'e' la tua lingua madre?
(What is your preferred language?)`,
	},
	MESSAGE_TEXT_CHOOSE_UI_LANGUAGE: {
		"ru-RU": "На каком языке вы хотели бы общаться со мной?",
		"en-US": "Which language you would like to talk to me?",
		"fa-IR": "دوست دارید به چه زبانی با من صحبت کنید؟",
		"it-IT": "In quale lingua preferisci parlarmi?",
	},
	MESSAGE_TEXT_LOCALE_CHANGED: {
		"ru-RU": "Вы поменяли язык на %v",
		"en-US": "You've switched language to %v",
		"fa-IR": "شما زبان را به %v تغییر دادید. ",
		"it-IT": "Hai cambiato lingua in %v",
	},
	MESSAGE_TEXT_WHATS_NEXT: {
		"ru-RU": "Что будем делать дальше?",
		"en-US": "What's next?",
		"fa-IR": "اقدام بعدی چیست؟",
		"it-IT": "Cosa posso fare ora?",
	},
	MESSAGE_TEXT_WHATS_NEXT_HINT: {
		"ru-RU": `
Если вы дали в долг воспользуйтесь командой /дал.
Если вы одолжили что-то - командой /взял.

Или воспользуйтесь меню внизу экрана.
`,
		"en-US": `
If you borrowed from someone to record it use /got.
If you lent to someone to record it use /gave.

Or use menu at the bottom.
`,

		"fa-IR": `
اگر از کسی قرض گرفته اید برای ثبت آن از /قرض_گرفتن استفاده کنید.
اگر به کسی قرض داده اید برای ثبت آن از /قرض_دادن استفاده کنید.

یا از منوی پایین استفاده نمایید.`,

		"it-IT": `
Se qualcuno ti ha prestato qualcosa per memorizzarlo usa /got.
Se hai prestato qualcosa a qualcuno per memorizzarlo usa /gave.

O usa il menu' qui sotto.

`,
	},
	MESSAGE_TEXT_HISTORY_HEADER: {
		"ru-RU": "История",
		"en-US": "History",
		"fa-IR": "سوابق",
		"it-IT": "Cronologia",
	},
	MESSAGE_TEXT_HISTORY_NO_RECORDS: {
		"ru-RU": "У вас пока нет ни одной записи.",
		"en-US": "You don't have any records yet.",
		"fa-IR": "شما هنوز هیچ ثبت سابقه ای ندارید.",
		"it-IT": "Non hai nulla memorizzato.",
	},
	MESSAGE_TEXT_HISTORY_LIST: {
		"ru-RU": `<b>%v</b> <i>(%d последних)</i>
─────────────
%v`,

		"en-US": `<b>%v</b> <i>(last %d):</i>
─────────────
%v`,

		"fa-IR": `<b>%v</b> <i>(آخرین %d):</i>
─────────────
%v`,

		"it-IT": `<b>%v</b> <i>(ultimi %d):</i>

─────────────
%v`,
	},
	MESSAGE_TEXT_BALANCE_IS_ZERO: {
		"ru-RU": "Нет записей о текущих долгах.",
		"en-US": "You have no records on current debts.",
		"fa-IR": "شما در خصوص بدهی های اخیر ثبت سابقه ای ندارید.",
		"it-IT": "Non hai nulla memorizzato nel debito corrente.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_TOTAL_INTRO: {
		"ru-RU": "Всего",
		"en-US": "Total",
		"fa-IR": "مجموع",
		"it-IT": "Totale",
	},
	MESSAGE_TEXT_PRIMARY_CURRENCY_IS_SET_TO: {
		"ru-RU": "OK, теперь я буду использовать '%v' как основную валюту.",
		"en-US": "OK, from now on I will use '%v' as a primary currency.",
		"fa-IR": "بسیار خوب، از الان من از '%v' بعنوان واحد پول اولیه استفاده می کنم",
		"it-IT": "OK, da ora in poi usero' '%v' come moneta principale.",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_TO_USER: {
		"ru-RU": "%v - долг вам %v",
		"en-US": "%v - owes you %v",
		"fa-IR": "%v - %v به شما بدهکار است ",
		"it-IT": "%v - ti deve %v.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_TO_USER: {
		"ru-RU": "Вам должны %v",
		"en-US": "Owes to you %v",
		"fa-IR": "%v به شما بدهکار است ",
		"it-IT": "Hai un credito di %v",
	},
	MESSAGE_TEXT_ON_RETURN_USER_DOES_NOT_OWE_ANYTHING_TO_COUNTERPARTY_ANYMORE: {
		"ru-RU": "Поздравляем! У вас не осталось долгов перед <b>%v</b>.",
		"en-US": "Congratulations! You don't owe anything more to <b>%v</b>.",
		"fa-IR": "تبریک! شما دیگر چیزی به <b>%v</b> بدهکار نیستید .",
		"it-IT": "Bravo! Hai saldato il tuo debito con <b>%v</b>.",
	},
	MESSAGE_TEXT_ON_RETURN_COUNTERPARTY_DOES_NOT_OWE_ANYTHING_TO_USER_ANYMORE: {
		"ru-RU": "У <b>%v</b> больше не осталось долгов перед вами.",
		"en-US": "<b>%v</b> does not owe anything more to you.",
		"fa-IR": "<b>%v</b> دیگر چیزی به شما بدهکار نیست",
		"it-IT": "<b>%v</b> ha saldato il suo debito verso di te.",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_BY_USER: {
		"ru-RU": "Вы должны %v",
		"en-US": "You owe %v",
		"fa-IR": "شما %v بدهکار هستید ",
		"it-IT": "Hai un debito di %v",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_BY_USER: {
		"ru-RU": "%v - вы должны %v",
		"en-US": "%v - you owe %v",
		"fa-IR": "%v - شما %v بدهکار هستید ",
		"it-IT": "%v - tu gli/le devi %v",
	},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY: {
		"ru-RU": "Какая валюта для вас основная?",
		"en-US": "What is your primary currency?",
		"fa-IR": "واحد پولی اولیه شما چیست؟",
		"it-IT": "Qual'e' la tua valuta principale?",
	},
	MESSAGE_TEXT_FAILED_TO_DELETE_USER: {
		"ru-RU": "Не удалось удалить данные пользователя: %v",
		"en-US": "Failed to delete user: %v",
		"fa-IR": "حذف کاربر ناموفق بود: %v",
		"it-IT": "Errore durante la cancellazione dell'utente: %v",
	},
	MESSAGE_TEXT_USER_DELETED: {
		"ru-RU": "Данные пользователя удалены",
		"en-US": "User's data has been deleted",
		"fa-IR": "اطلاعات کاربر حذف شد.",
		"it-IT": "I dati dell'utente sono stati cancellati",
	},
	MESSAGE_TEXT_PLEASE_WAIT_WHILE_WE_GENERATE_INVITE_CODE: {
		"ru-RU": "Пожалуйста подождите пока мы генерируем секретный код доступа...",
		"en-US": "Please wait a moment while we are generating a security access code...",
		"fa-IR": "لطفاً کمی صبر کنید تا ما یک کد دسترسی امنیتی  ایجاد کنیم.",
		"it-IT": "Aspetta un attimo mentre sto generando un codice di accesso sicuro...",
	},
	MESSAGE_TEXT_RETURN_ASK_TO_CHOOSE_COUNTERPARTY: {
		"ru-RU": "Выберете кому вы вернули долг или кто вернул его вам.",
		"en-US": "Please choose who returned the debt or to who you returned it.",
		"fa-IR": "لطفاً انتخاب کنید چه کسی بدهی اش را به شما پرداخت کرده یا شما بدهیتان را به چه کسی بازپرداخت نموده اید.",
		"it-IT": "Scegli con chi hai sanato un credito o un debito.",
	},
	MESSAGE_TEXT_CHOOSE_DEBT_THAT_HAS_BEEN_RETURNED: {
		"ru-RU": "Выберите долг который был возвращён целиком или частично.",
		"en-US": "Please choose a debt that has been returned fully or partially.",
		"fa-IR": "لطفاً انتخاب کنید تمام یا بخشی از کدام بدهی پرداخت شده است.",
		"it-IT": "Scegli un debito che e' stato restituito completamente o parzialmente.",
	},
	MESSAGE_TEXT_PLEASE_ACKNOWLEDGE_TRANSFER: {
		"ru-RU": "Пожалуйста подтвердите или отклоните эту транзакцию.",
		"en-US": "Please confirm or decline this transfer.",
		"fa-IR": "لطفاً این تراکنش را تایید یا رد نمایید.",
		"it-IT": "Conferma o rifiuta questo debito/credito.",
	},
	MESSAGE_TEXT_ALREADY_ACCEPTED_TRANSFER: {
		"ru-RU": "Эта транзакция уже подтверждена.",
		"en-US": "This transfer has been accepted already.",
		"fa-IR": "این تراکنش قبلا قبول شده است.",
		"it-IT": "Questo debito/credito e' gia' stato accettato.",
	},
	MESSAGE_TEXT_ALREADY_DECLINED_TRANSFER: {
		"ru-RU": "Эта транзакция уже отклонена.",
		"en-US": "This transfer has been declined already.",
		"fa-IR": "این تراکنش قبلاً رد شده است.",
		"it-IT": "Questo debito/credito e' gia' stato rifiutato.",
	},
	MESSAGE_TEXT_RECEIPT_LINK: {
		"ru-RU": "Подробнее тут: %v",
		"en-US": "Details here: %v",
		"it-IT": "Maggiori dettagli qui: %v",
		"fa-IR": "جزئیات: %v",
	},
	MESSAGE_TEXT_ASK_PHONE_NUMBER_OF_COUNTERPARTY: {
		"ru-RU": "Пожалуйста напишите номер телефона <b>%v</b>.",
		"en-US": "Please provide phone number for <b>%v</b>",
		"it-IT": "Per favore fornisci il numero di telefono di <b>%v</b>",
		"fa-IR": "لطفا شماره تلفن ایشان را وارد کنید <b>%v</b>",
	},
	MESSAGE_TEXT_USE_CONTACT_TO_SEND_PHONE_NUMBER: {
		"ru-RU": "Если номер телефона есть в записной книжке <b>воспользуйтесь кнопкой %v</b> (скрепка) чтобы отправить контакт.",
		"en-US": "If phone number is in your address book you can <b>use %v button</b> to send the contact.",
		"fa-IR": "اگر شماره تلفن در فهرست مخاطبین شما وجود دارد شما می توانید <b> با استفاده از این %v دکمه</b> تماس را ارسال نمایید.",
		"it-IT": "Se il numero e' nella tua rubrica, puoi <b> usare il pulsante %v</b> per inviare il contatto.",
	},
	MESSAGE_TEXT_ABOUT_PHONE_NUMBER_FORMAT: {
		"ru-RU": "Номер должен быть в международном формате:\n\t* Начинаться со знака '+' и кода страны\n\t* Состоять только из цифр\nПример: <pre>+</pre><b>7</b><code>999012345678</code>",
		"en-US": "The number should be in international standard:\n\t* Starts with '+' following by country code\n\t* Consist of numbers only\nExample: <pre>+</pre><b>1</b><code>999012345678</code>",
		"fa-IR": "شماره باید به صورت استاندارد بین المللی باشد\n\t* با '+' شروع شده و بدنبال آن کد کشور وارد شود\n\t* تنها شامل اعداد باشد\nمثال: <pre>+</pre><b>1</b><code>999012345678</code>",
		"it-IT": "Il numero deve essere in formato internazionale:\n\t* Inizia con '+' seguito dal codice del paese (Italia +39)\n\t* \nEsempio: <pre>+</pre><b>39</b><code>34612345678</code>",
	},
	MESSAGE_TEXT_THIS_NUMBER_WILL_BE_USED_TO_SEND_RECEIPT: {
		"ru-RU": "На этот номер мы отправим SMS:",
		"en-US": "Will send an SMS to this number:",
		"fa-IR": "یک پیام کوتاه به این شماره ارسال خواهد شد:",
		"it-IT": "Invieremo un SMS a questo numero:",
	},
	MESSAGE_TEXT_COUNTERPARTY_OWES_YOU_SINGLE_DEBT: {
		"ru-RU": `<b>%v</b> одалживал(а) у вас <b>%v</b>.`,
		"en-US": `<b>%v</b> owed to you <b>%v</b>.`,
		"fa-IR": `<b>%v</b> به شما بدهکار بوده <b>%v</b>.`,
		"it-IT": `<b>%v</b> e' in debito di <b>%v</b>.`,
	},
	MESSAGE_TEXT_YOU_OWE_TO_COUNTERPARTY_SINGLE_DEBT: {
		"ru-RU": `<b>%v</b> одалживал(а) вам <b>%v</b>.`,
		"en-US": "You owe to <b>%v</b> <b>%v</b>.",
		"fa-IR": "شما بدهکار هستید به <b>%v</b> <b>%v</b>.",
		"it-IT": `Tu devi dare a <b>%v</b> <b>%v</b>.`,
	},
	MESSAGE_TEXT_IS_IT_RETURNED_IN_FULL: {
		"ru-RU": `Возвращено полностью?

		<i>Если частично можете сразу написать сумму.</i>`,

		"en-US": `Has this debt been returned in full?

		<i>If partially you can enter amount right away.</i>`,

		"fa-IR": `آیا این بدهی بصورت کامل بازپرداخت شده است؟

		<i>اگر بخشی از بدهی پرداخت شده است شما میتوانید مبلغ را وارد کنید.</i>`,

		"it-IT": `Il debito e' stato saldato?

		<i>Se la risposta e' NO puoi inserire l'ammontare da sottrarre, direttamente qui sotto.</i>`,
	},
	MESSAGE_TEXT_PLEASE_HELP_MAKE_IT_BETTER: {
		"ru-RU": `Эта программа <b>бесплатна</b>. <a href="https://debtstracker.io/ru/help-us">Помогите</a> сделать её лучше!`,
		"en-US": `This program is <b>free to use</b>. Please <a href="https://debtstracker.io/en/help-us">help</a> to make it better!`,
		"it-IT": `Questo programma e' <b> completamente gratis</b>. Per favore <a href="https://debtstracker.io/">aiuta</a> a migliorarlo!`,
		"fa-IR": `این برنامه <b>رایگان می باشد</b>.لطفاً <a href="https://debtstracker.io/">به ما کمک کنید</a>تا آنرا بهبود دهیم!`,
	},
	BUTTON_TEXT_YOU_OWE_AMOUNT_TO_SOMEONE: {
		"ru-RU": "%v | вы должны: %v",
		"en-US": "%v | you owe: %v",
		"fa-IR": "%v | شما بدهکارید: %v",
		"it-IT": "%v | tu devi: %v",
	},
	BUTTON_TEXT_SOMEONE_OWES_TO_YOU_AMOUNT: {
		"ru-RU": "%v | долг вам: %v",
		"en-US": "%v | owes to you: %v",
		"fa-IR": "%v | به شما بدهکار است: %v",
		"it-IT": "%v | ti deve: %v",
	},
	BUTTON_TEXT_DEBT_RETURNED_FULLY: {
		"ru-RU": "Да, целиком",
		"en-US": "Yes, fully",
		"fa-IR": "بله، به صورت کامل",
		"it-IT": "Si, completamente",
	},
	BUTTON_TEXT_DEBT_RETURNED_PARTIALLY: {
		"ru-RU": "Нет, только часть",
		"en-US": "No, just partially",
		"fa-IR": "خیر، تنها قسمتی",
		"it-IT": "No, parzialmente",
	},
	MESSAGE_TEXT_ATTEMPT_TO_USE_OWN_INVITE: {
		"ru-RU": "Хорошая попытка пригласить самого себя ;)",
		"en-US": "You should not use your own invite ;)",
		"fa-IR": "نباید از دعوت خود استفاده کنید ;)",
		"it-IT": "Non dovresti usare il tuo codice invito con te stesso :)",
	},
	MESSAGE_TEXT_WELCOME_ONBOARDING_INVITE_ACCEPTED: {
		"ru-RU": "Спасибо за то что воспользовались приглашением!",
		"en-US": "Welcome and thanks for accepting the invite!",
		"fa-IR": "خوش آمدید و سپاسگزاریم که دعوت را پذیرفتید!",
		"it-IT": "Benvenuto e grazie per aver accettato l'invito!",
	},
	MESSAGE_TEXT_FOR_COUNTERPARTY_ONLY: {
		"ru-RU": "Это действие доступно только для %v",
		"en-US": "This action for %v only.",
		"fa-IR": "این عمل تنها برای %v می باشد.",
		"it-IT": "Questa azione e' disponibile solo per %v.",
	},
	BUTTON_TEXT_SEE_RECEIPT_DETAILS: {
		"ru-RU": "Показать детали",
		"en-US": "Show receipt details",
		"fa-IR": "جزئیات رسید را نشان بده",
		"it-IT": "Mostra i dettagli del debito/credito",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_EMAIL: {
		"ru-RU": "Вы решили пригласить друга через email.",
		"en-US": "You've selected to invite friend by email.",
		"fa-IR": "شما انتخاب کردید که یک دوست را بوسیله ایمیل دعوت کنید.",
		"it-IT": "Hai scelto di invitare l'amico tramite email.",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_SMS: {
		"ru-RU": "Вы решили пригласить друга через SMS.",
		"en-US": "You've selected to invite friend by SMS.",
		"fa-IR": "شما انتخاب کردید که یک دوست را بوسیله پیام کوتاه دعوت کنید",
		"it-IT": "Hai scelto di invitare l'amico tramite SMS.",
	},
	MESSAGE_TEXT_ABOUT_INVITES: {
		"ru-RU": `На данный момент доступ к нашему боту ограничен, но вы можете пригласить друга.

Как вы хотите передать код приглашение?`,

		"en-US": `At the moment access to our bot is limited but you can invite your friend.

How do you want to pass the invite code?`,

		"fa-IR": `در حال حاضر دسترسی به ربات محدود می باشد ولی شما می توانید دوست خود را دعوت کنید.

How do آیا میخواهید کد دعوت را ارسال کنید؟`,

		"it-IT": `AL momento l'accesso al nostro bot e' limitato ma puoi comunque invitare gli amici.

Come vuoi inviargli il codice invito?`,
	},
	MESSAGE_TEXT_USER_BLOCKED_TRANSFER_NOTIFICATIONS_BY: {
		"ru-RU": "%v заблокировал получение оповешений о транзакиях через: %v.",
		"en-US": "%v blocked reminders about transactions by: %v",
		"fa-IR": "%v یادآور تراکنش مسدود شده است بوسیله ی: %v",
		"it-IT": "%v bloccato promemoria riguardo la transazione da: %v.",
	},
	COMMAND_TEXT_WAIT_A_SECOND: {
		"ru-RU": "Секундочку...",
		"en-US": "Wait a second...",
		"fa-IR": "یک ثانیه صبر کنید ...",
		"it-IT": "Solo un attimo...",
	},
	HTML_USING_TELEGRAM: {
		"ru-RU": "используя Telegram мессенджер",
		"en-US": "using Telegram messenger",
		"fa-IR": "استفاده از پیام رسان تلگرام",
		"it-IT": "usa Telegram",
	},
	COMMAND_TEXT_ACCEPT: {
		"ru-RU": "Согласиться",
		"en-US": "Accept",
		"fa-IR": "قبول",
		"it-IT": "Accetta",
	},
	//BUTTON_TEXT_ACCEPT_TRANSFER_USING_TELEGRAM:{
	//	"ru-RU": "Подтвердить ",
	//	"en-US": "Accept ",

	//	"fa-IR": "قبول ",

	//  "it-IT": "Accetta",

	//},
	//BUTTON_TEXT_DECLINE_TRANSFER_USING_TELEGRAM:{
	//	"ru-RU": "Отказаться (используя Telegram)",
	//	"en-US": "Decline (using Telegram messenger)",

	//	"fa-IR": "رد ( استفاده از پیام رسان تلگرام)",

	//  "it-IT": "Rifiuta (usando Telegram)",

	//},
	COMMAND_TEXT_DECLINE: {
		"ru-RU": "Отклонить",
		"en-US": "Decline",
		"fa-IR": "رد",
		"it-IT": "Rifiuta",
	},
	COMMAND_TEXT_ACCEPT_INVITE: {
		"ru-RU": "Принять приглашение",
		"en-US": "Accept invite",
		"fa-IR": "قبول دعوت",
		"it-IT": "Accetta l'invito",
	},
	COMMAND_TEXT_VIEW_RECEIPT_DETAILS: {
		"ru-RU": "Посмотреть квитанцию",
		"en-US": "See receipt details",
		"fa-IR": "دیدن جزئیات رسید",
		"it-IT": "Vedi dettagli",
	},
	COMMAND_TEXT_OTHER_WAYS_TO_SEND_INVITE: {
		"ru-RU": "Другие способы отправить приглашение",
		"en-US": "Other ways to send an invite",
		"fa-IR": "سایر راههای ارسال دعوت",
		"it-IT": "Altri modi per inviare un invito",
	},
	COMMAND_TEXT_SEND_MY_PHONE_NUMBER: {
		"ru-RU": "Отправить мой номер",
		"en-US": "Send my phone number",
		"fa-IR": "شماره تلفن مرا ارسال کنید",
		"it-IT": "Invia il mio numero",
	},
	COMMAND_TEXT_SEND_BY_EMAIL: {
		"ru-RU": "Через Email",
		"en-US": "By Email",
		"fa-IR": "بوسیله ی ایمیل",
		"it-IT": "Tramite email",
	},
	COMMAND_TEXT_SEND_BY_SMS: {
		"ru-RU": "Через SMS",
		"en-US": "By SMS",
		"it-IT": "Tramite SMS",
		"fa-IR": "بوسیله پیام کوتاه",
	},
	COMMAND_TEXT_INVITE_BY_TELEGRAM: {
		"ru-RU": "Пригласить через Telegram",
		"en-US": "Invite By Telegram",
		"it-IT": "Tramite Telegram",
		"fa-IR": "دعوت با تلگرام",
	},
	MESSAGE_TEXT_INVITE_CREATED: {
		"ru-RU": `Мы отправили код приглашения на указынный вами адрес. (#%v)

Когда ваш друг потдвердит приглашение у вас будут общий баланс и транзакции (только между вами). Это поможет вам минимизировать усилия по ведению учёта.`,

		"en-US": `We've sent an invite code to your friend. (#%v)

Once your friend accepts invitation you'll share balance & transfers between you to make sure you both are on the same page with minimum efforts.`,
		"it-IT": `Abbiamo inviato il codice invito al tuo amico. (#%v)

Una volta che il tuo amico accetta l'invito potrete condividere i bilanci ed i trasferimenti con il minimo sforzo.`,
		"fa-IR": `ما برای دوست شما یک  پیام دعوت ارسال کردیم. (#%v)

وقتی دوست شما دعوت را بپذیرد شما تراز و مبادلات بین خود را به اشتراک می گذارید تا با کمترین تلاش از درک مشترک میان خود اطمینان حاصل کنید. `,
	},
	MESSAGE_TEXT_INVITE_BY_EMAIL: {
		"ru-RU": "Введите email вашего друга на который мы отправим код приглашения.",
		"en-US": "Please enter emaill address of your friend where we should send an invite code.",
		"it-IT": "Inserisci l'email dell'amico al quale inviare il codice invito.",
		"fa-IR": "لطفاً آدرس ایمیل دوست خود را وارد کنید تا ما از آن طریق کد دعوت را ارسال نماییم.",
	},
	MESSAGE_TEXT_INVITE_ASK_EMAIL_FOR_RECEIPT: {
		"ru-RU": "Введите email вашего друга (%v) на который мы отправим квитанцию подтверждения.",
		"en-US": "Please enter emaill address of your friend (%v) where we should send the receipt.",
		"fa-IR": "لطفا آدرس ایمیل دوست خود را وارد کنید (%v) تا ما از آن طریق کد دعوت را ارسال نماییم.",
		"it-IT": "Inserisci l'email del tuo amico (%v) alla quale potremo inviare il debito/credito",
	},
	MESSAGE_TEXT_INVITE_BY_SMS: {
		"ru-RU": "Введите номер телефона вашего друга на который мы отправим код приглашения.",
		"en-US": "Please share a contact or enter phone number of your friend where we should send an invite code.",
		"it-IT": "COndividi il contatto o inserisci il numero di telefono del tuo amico al quale invieremo il codice invito.",
		"fa-IR": "لطفا شماره دوستان را که می خواهید برای او کد دعوت بفرستیم از لیست مخاطبین به اشتراک گذاشته یا آن را وارد کنید.",
	},
	MESSAGE_TEXT_INVITE_BY_TELEGRAM: {
		"ru-RU": "Вставьте в чат контакт вашего друга которому вы хотите отправить приглашение.",
		"en-US": "Please share a contact of your friend you wish to send an invite code:",
		"it-IT": "Condividi il contatto di un amico al quale desideri inviare il codice invito.",
		"fa-IR": "لطفا اطلاعات تماس دوستتان را که میخواهید برایشان کد دعوت ارسال شود به اشتراک بگذارید.",
	},
	MESSAGE_TEXT_INVALID_EMAIL: {
		"ru-RU": "Неверный email. Проверьте и попробуйте ещё раз? /menu",
		"en-US": "Invalid email. Check and try it again? /menu",
		"fa-IR": "ایمیل غیر معتبر می باشد. آیا بررسی نموده، دوباره سعی می کنید؟ /منو",
		"it-IT": "Email scorretta. COntrolla e riprova. /menu",
	},
	MESSAGE_TEXT_INVALID_YEAR: {
		"ru-RU": "Неверный год. Год должен быть 2 или 4 цифры (<i>например 2016 или 16)</i>).",
		"en-US": "Invalid year. The year part should be 2 or 4 numbers (<i>e.g. 2016 or 16</i>).",
		"fa-IR": "سال غیرمعتبر می باشد. سال باید به صورت 2 یا 4 رقمی وارد شود (<i>برای مثال 16 یا 2016</i>).",
		"it-IT": "Anno scorretto. L'anno dev;essere composta da 2 o 4 numeri (<i>esempio 2017 oppure 17</i>)",
	},
	MESSAGE_TEXT_INVALID_MONTH: {
		"ru-RU": "Неверный месяц. Месяц должен быть задан целым числом от 1 до 12.",
		"en-US": "Invalid month. Month should be an integer from 1 to 12.",
		"fa-IR": "ماه غیر معتبر می باشد. ماه باید به صورت عددی صحیح بین 1 تا 12 باشد.",
		"it-IT": "Mese scorretto. Il mese dovrebbe essere un numero da 1 a 12.",
	},
	MESSAGE_TEXT_INVALID_DAY: {
		"ru-RU": "Неверный день. День должен быть задан целым числом от 1 до 31.",
		"en-US": "Invalid day. The day should be an integer from 1 to 31.",
		"fa-IR": "روز غیر معتبر می باشد. روز باید عددی صحیح بین 1 تا 31 باشد.",
		"it-IT": "Giorno scorretto. Il giorno dovrebbe essere un numero da 1 a 31.",
	},
	MESSAGE_TEXT_INVALID_DATE: {
		"ru-RU": "Неверный формат даты. Например для 20 февраля 2019 года надо ввести: 20.02.2019 или 20.02.19",
		"en-US": "Invalid date format. For exampel for 20 February 2019 please submit: 20.02.2019 or 20.02.19",
		"fa-IR": "فرمت تاریخ غیر معتبر می باشد. برای مثال برای 20 فوریه 2019 لطفا اینگونه وارد کنید: 20.02.2019 یا 20.02.19",
		"it-IT": "Formato data sbagliato. Esempio: per il 20 Febbraio 2019 inserisci: 20.02.2019 oppure 20.02.19",
	},
	MESSAGE_TEXT_INVALID_PHONE_NUMBER: {
		"ru-RU": "Неверный номер. Проверьте и попробуйте ещё раз? /menu",
		"en-US": "Invalid phone number. Check and try it again? /menu",
		"fa-IR": "شماره تلفن غیر معتبر می باشد. آیا بررسی نموده، مجدداً سعی می کنید؟ /منو",
		"it-IT": "Numero di telefono invalido. Controlla e riprova. /menu",
	},
	MESSAGE_TEXT_PHONE_NUMBER_IS_NOT_SMS_CAPABLE: {
		"ru-RU": "Данный номер не принимает SMS. Попробуйте другой номер? /menu",
		"en-US": "This phone number not able to receive SMS. Try another number? /menu",
		"fa-IR": "این شماره تلفن قادر به دریافت پیام کوتاه نمی باشد. آیا شماره دیگری را امتحان میکنید؟ /منو",
		"it-IT": "Questo numero di telefono non e' abilitato a ricevere SMS. Vuoi provare un altro numero? /menu",
	},
	MESSAGE_TEXT_NO_CONTACT_RECEIVED: {
		"ru-RU": "Мы не получили контакта. Тут инструкция как это сделать. /menu",
		"en-US": "We have not received any contacts. INSTRUCTION HOW TO DO IT. /menu",
		"fa-IR": "ما هیچ اطلاعات تماسی دریافت نکردیم. دستورالعمل چگونگی انجام این کار. /منو",
		"it-IT": "Non abbiamo ricevuto nesusn contatto. ISTRUZIONI SU COME FARE. /menu",
	},
	MESSAGE_TEXT_CONTACT_NAME_IS_NUMBER: {
		"ru-RU": "Вы ввели только цифры в качестве имени контакта. Пожалуйста используйте текстовые символы.",
		"en-US": "You've entered just digits for a contact name. Please use some text characters.",
		"fa-IR": "شما تنها اعداد را برای نام مخاطب وارد کرده اید. لطفا کاراکتر های متنی وارد کنید.",
		"it-IT": "Hai inserito solamente numeri per un nome contatto. Usa anche alcune lettere.",
	},
	MESSAGE_TEXT_CURRENCY_NAME_IS_NUMBER: {
		"ru-RU": "Вы ввели только цифры в качестве номинала. Пожалуйста используйте текстовые символы.",
		"en-US": "You've entered just digits for currency. Please use some text characters.",
		"fa-IR": "شما تنها اعداد را برای واحد پولی وارد کرده اید. لطفا کاراکترهای متنی وارد کنید.",
		"it-IT": "Hai inserito solamente numeri per la valuta. Usa anche alcune lettere.",
	},
	MESSAGE_TEXT_HISTORY_ROW_TO_USER_WITH_NAME: {
		"ru-RU": "%v - %s ⇒ Вам : %s",
		"en-US": "%v - %s ⇒ to you: %s",
		"fa-IR": "%v - %s ⇒ به شما: %s",
		"it-IT": "%v - %s ⇒ a te: %s",
	},
	MESSAGE_TEXT_HISTORY_ROW_FROM_USER_WITH_NAME: {
		"ru-RU": "%v - Вы ⇒ %s : %s",
		"en-US": "%v - You ⇒ %s : %s",
		"fa-IR": "%v - شما ⇒ %s : %s",
		"it-IT": "%v - Tu ⇒ %s : %s",
	},
	MESSAGE_TEXT_LETS_SEND_SMS: {
		"ru-RU": "Давайте отправим SMS",
		"en-US": "Let's send SMS",
		"fa-IR": "پیام کوتاه ارسال کنید",
		"it-IT": "Inviamo un SMS",
	},
	MESSAGE_TEXT_SMS_QUEUING_FOR_SENDING: {
		"ru-RU": "SMS ставится в очередь на отправку на номер %v...",
		"en-US": "Queuing SMS for sending to number %v...",
		"fa-IR": "پیام کوتاه برای ارسال به شماره مقابل در حال قرارگیری در نوبت ارسال می باشد %v...",
		"it-IT": "SMS in coda per l'invio al numero %v...",
	},
	MESSAGE_TEXT_SMS_QUEUED_FOR_SENDING: {
		"ru-RU": "SMS поставлена в очередь на отправку на номер %v",
		"en-US": "SMS is queued for sending to number %v",
		"fa-IR": "پیام کوتاه برای شماره مقابل در نوبت ارسال قرار گرفت %v",
		"it-IT": "SMS inserito in coda per l'invio al numero %v",
	},
	MESSAGE_TEXT_BALANCE_HEADER: {
		"ru-RU": "Баланс",
		"en-US": "Balance",
		"fa-IR": "تراز",
		"it-IT": "Bilancio",
	},
	MESSAGE_TEXT_RECEIPT_AVAILABLE_CHANNELS: {
		"ru-RU": "Извините, в данный момент доступны только эти каналы для отправки квитанции:",
		"en-US": "Sorry, at the moment just this channels are available for sending a receipt:",
		"fa-IR": "متاسفانه، در حال حاضر تنها این کانالها برای ارسال رسید در دسترس می باشند.",
		"it-IT": "Spiacenti ma al momento solo questi canali sono disponibili per inviare debiti/crediti:",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_TELEGRAM: {
		"ru-RU": "Квитанция отправлена через телеграм.",
		"en-US": "Receipt sent through Telegram.",
		"fa-IR": "رسید از طریق تلگرام ارسال شد.",
		"it-IT": "Notifica inviata tramite Telegram",
	},
	MESSAGE_TEXT_RECEIPT_NOT_SENT_AS_COUNTERPARTY_HAS_DISABLED_TG_BOT: {
		"ru-RU": "Квитанция НЕ отправлена через телеграм так как %v удалил чат с ботом.",
		"en-US": "Receipt NOT sent through Telegram as %v has deleted chat with the bot.",
		"fa-IR": "از آنجایی که %v چت انجام شده با روبات را حذف کرده است رسید از طریق تلگرام ارسال نشد.",
		"it-IT": "Notifica NON inviata tramite Telegram a %v perche' ha cancellato la chat con il bot",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_EMAIL: {
		"ru-RU": "Квитанция отправлена через email. (id: %v)",
		"en-US": "Receipt sent through email. (id: %v)",
		"fa-IR": "رسید از طریق ایمیل ارسال شد. (id: %v)",
		"it-IT": "Notifica inviata tramite email (id: %v)",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_SMS: {
		"ru-RU": "Квитанция отправлена через SMS.",
		"en-US": "Receipt sent through SMS.",
		"fa-IR": "رسید از طریق پیام کوتاه ارسال شد.",
		"it-IT": "Notifica inviata tramite SMS",
	},
	MESSAGE_TEXT_SWITCH_TO_PM_TO_VIEW_RECEIPT: {
		"ru-RU": "Переключитьсь на чат с ботом чтобы посмотреть квитанцию",
		"en-US": "Switch to private mode to see receipt details.",
		"fa-IR": "انتقال به حالت خصوصی جهت رویت جزئیات رسید.",
		"it-IT": "Passa alla modalita' privata per vedere i dettagli dei tuoi crediti/debiti.",
	},
	MESSAGE_TEXT_RECEIPT_VIEWED_BY_COUNTERPARTY: {
		"ru-RU": "Квитанция просмотрена",
		"en-US": "Receipt viewed",
		"fa-IR": "رسید رویت شد.",
		"it-IT": "Debiti visti",
	},
	MESSAGE_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"ru-RU": "Вы можете посмотреть свой номер телефона в ожидаемоем нами формате.",
		"en-US": "You can view your own phone number in the format we are expecting.",
		"fa-IR": "شما می توانید شماره تلفن خود را با فرمتی که ما انتظار داریم ببینید.",
		"it-IT": "Puoi visionare il tuo numero di telefono nel formato previsto.",
	},
	COMMAND_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"ru-RU": "Посмотреть мой номер в ожидаемоем формате",
		"en-US": "View my number in the expectd format",
		"fa-IR": "رویت شماره خودم با فرمت مورد انتظار",
		"it-IT": "Mostra il mio numero nel formato previsto",
	},
	INLINE_BUTTON_SHOW_FULL_HISTORY: {
		"ru-RU": "Показать всю историю",
		"en-US": "Show full history",
		"fa-IR": "نمایش کامل سوابق",
		"it-IT": "Mostra cronologia completa",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_NOT_A_NUMBER: {
		"ru-RU": "Это не цифра",
		"en-US": "it is not a number",
		"fa-IR": "این یک شماره نیست",
		"it-IT": "Non e' un numero",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_IS_NEGATIVE: {
		"ru-RU": "Цифра должна быть положительной (<i>больше нуля</i>)",
		"en-US": "The number should be positive (<i>greater than 0</i>)",
		"fa-IR": "شماره باید مثبت باشد (<i>بزرگتر از 0</i>)",
		"it-IT": "Il numero dovrebbe essere positivo (<i>maggiore di 0<i/>)",
	},
	MESSAGE_TEXT_ASK_HOW_MUCH_HAS_BEEN_RETURNED: {
		"ru-RU": "Сколько было возвращено?",
		"en-US": "How much have been returned?",
		"fa-IR": "چه مقدار بازپرداخت شده است؟",
		"it-IT": "Quanto ti e' stato restituito?",
	},
	MESSAGE_TEXT_HELP: {
		"ru-RU": "Вы можете сообщить о проблеме или сделать предложения по улучшению программы на нашем сайте.",
		"en-US": "Please report any issue or submit a feature request at our website.",
		"fa-IR": "لطفاً در وب سایت ما هرگونه مسئله ای را اعلام فرموده یا درخواست خود را مطرح نمایید.",
		"it-IT": "Segnala un problema o proponi un miglioramento sul nostro sito web.",
	},
	COMMAND_TEXT_OPEN_USER_REPORT: {
		"ru-RU": "Cтраница поддержки ",
		"en-US": "Support page",
		"fa-IR": "صفحه پشتیبانی",
		"it-IT": "Pagina d'aiuto",
	},
	COMMAND_TEXT_REPORT_A_BUG: {
		"ru-RU": "Сообщить об ошибке",
		"en-US": "Report a bug",
		"fa-IR": "گزارش یک باگ",
		"it-IT": "Segnala un bug",
	},
	COMMAND_TEXT_SUBMIT_AN_IDEA: {
		"ru-RU": "Предложить идею",
		"en-US": "Add an idea",
		"fa-IR": "یک ایده اضافه کنید.",
		"it-IT": "Proponi un idea",
	},
	MESSAGE_TEXT_WELCOME: {
		"ru-RU": `Привет, я Коллектиус - Ваш персональный счетовод и коллектор.

Могу записать кто кому чего должен и, и при необходимости, напомнить когда должок пора возвращать.

Чего изволит новый хозяин?`,
		"en-US": `Hi, I'm Collectius - your personal accountant & collector.

I can record who is owing to whom and remind when the return is due.

What would you like to do?`,

		"fa-IR": `سلام، من کالکتیوس هستم - حسابدار شخصی و مامور وصول شما

من میتوانم اینکه چه کسی به چه کسی بدهکار است را ثبت کرده و زمان بازپرداخت را یادآوری کنم.

دوست دارید چکار کنید؟`,

		"it-IT": `Ciao, sono Collectius - il tuo contabile ed esattore.

Posso annotare chi deve soldi a chi e ricordarti la data di scadenza.

Cosa vorresti fare ora?`,
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE: {
		"ru-RU": "Хочу получить приглашение",
		"en-US": "I want to get an invite",
		"fa-IR": "می خواهم یک دعوت دریافت کنم.",
		"it-IT": "Voglio ottenere un invito",
	},
	COMMAND_TEXT_I_HAVE_INVITE: {
		"ru-RU": "У меня есть код приглашения",
		"en-US": "I have the invitation code",
		"fa-IR": "من کد دعوت را دارم.",
		"it-IT": "Ho il codice invito",
	},
	COMMAND_TEXT_I_HAVE_NOT_GOT_EMAIL: {
		"ru-RU": "Я не получил письма на email",
		"en-US": "I have not got any emails",
		"fa-IR": "من ایمیلی دریافت نکردم",
		"it-IT": "Non ho ricevuto nessun email",
	},
	MESSAGE_TEXT_ONBOARDING_TELL_ABOUT_INVITES: {
		"ru-RU": `<b>%v</b>,

На данный момент наш бот доступен только тем кто получил приглашение от друзей.

Если у вас нет кода можете оставить свои контакты и мы вышлем вам приглашение как только наступит ваша очередь.

Мы высылаем по 10 кодов в день первоочередникам + 1 случайным образом.`,

		"en-US": `<b>%v</b>,

At the moment our bot is available just by invitation from friends.

If you have no code you can leave your contact details and we'll send you an invite as soon as your queue is due.

We send 10 invites per day to those in the head of the queue and 1 randomly.`,

		"fa-IR": `<b>%v</b>,

درحال حاضر ربات ما تنها با دریافت دعوت از دوستان در دسترس می باشد.

اگر شما کدی در اختیار ندارید می توانید اطلاعات تماس خود را برای من وارد نموده و من به محض اینکه نوبت شما فرارسید یک دعوتنامه برایتان ارسال می کنم.

ما روزانه 10 دعوتنامه برای نفرات اول لیست انتظار و همچنین یک دعوتنامه تصادفی ارسال میکنیم.`,

		"it-IT": `<b>%v</b>,

Al momento il nostro bot e' disponibile solo tramite invito da amici.

Se non hai un codice puoi lasciarci il tuo contatto e ti manderemo un invito non appena sara' il tuo turno.

Inviamo 10 inviti al giorno ai primi 10 della lista d'attesa ed 1 in modo casuale.`,
	},
	EMAIL_INVITE_SUBJ: {
		"ru-RU": "Приглашение от {{.FromName}} - код: {{.InviteCode}}",
		"en-US": "An invite from {{.FromName}} - code: {{.InviteCode}}",
		"fa-IR": "دعوت از طرف {{.FromName}} - کد: {{.InviteCode}}",
		"it-IT": "Hai ricevuto un codice invito da {{.FromName}} - codice: {{.InviteCode}}",
	},
	SMS_INVITE_TEXT: {
		"ru-RU": `Привет {{.ToName}}, {{.FromName}} рекомендует приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Код приглашения: {{.InviteCode}}`,

		"fa-IR": `سلام{{.ToName}}, {{.FromName}} شما را دعوت کرده تا برنامه ردیابی بدهی ها را امتحان کنید.- https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,

		"en-US": `Hi {{.ToName}}, {{.FromName}} is inviting you to try debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Your personal invitation code is: {{.InviteCode}}`,

		"it-IT": `Ciao {{.ToName}}, {{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,
	},
	EMAIL_INVITE_TEXT: {
		"ru-RU": `Привет {{.ToName}},

{{.FromName}} приглашает тебя попробовать приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Ваш код приглашения: {{.InviteCode}}`,

		"en-US": `Hi {{.ToName}},

{{.FromName}} is inviting you to use debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

You invitation code is: {{.InviteCode}}`,

		"fa-IR": `سلام{{.ToName}},

{{.FromName}} شما را دعوت کرده تا از برنامه ردیابی بدهی ها استفاده کنید.- https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,

		"it-IT": `Ciao {{.ToName}},

{{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,
	},
	EMAIL_INVITE_HTML: {
		"ru-RU": `<p>Привет {{.ToName}}, </P

<p>{{.FromName}} приглашает тебя <a href = "https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">попробовать приложение для учёта долгов</a>.</p>

<p>Ваш код приглашения: <b>{{.InviteCode}}</b></p>`,

		"en-US": `<p>Hi {{.ToName}}, </p>

<p>{{.FromName}} is inviting you to <a href = "https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">try debts tracking app</a>.</p>

<p>Your invitation code is: <b>{{.InviteCode}}</b></p>`,

		"fa-IR": `<p>سلام{{.ToName}},</p>

<p>{{.FromName}} п شما را دعوت کرده به <a href="https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}"> امتحان برنامه ردیابی بدهی ها</a>.</p>

<p>You invitation code is: <b>{{.InviteCode}}</b></p>`,

		"it-IT": `<p>Ciao {{.ToName}},</p>

<p>{{.FromName}} ti ha invitato a provare <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">debts tracking app</a>.</p>

<p>Il tuo codice di invito personale e': <b>{{.InviteCode}}</b></p>`,
	},
	EMAIL_RECEIPT_SUBJ: {
		"ru-RU": "Запись о долге - {{.FromName}}",
		"en-US": "Debt record - {{.FromName}}",
		"fa-IR": "سوابق بدهی - {{.FromName}}",
		"it-IT": "Debito - {{.FromName}}",
	},
	EMAIL_RECEIPT_BODY_TEXT: {
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
		"en-US": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		"fa-IR": "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
		"it-IT": "{{.FromName}} ha creato un debito: {{.ReceiptURL}} ",
	},
	MESSENGER_RECEIPT_TEXT: {
		"ru-RU": "Я создал(а) запись о долге, подробности тут - {{.ReceiptURL}}",
		"en-US": "I've created a debt record regards you, see details at {{.ReceiptURL}}",
		"fa-IR": "من یک سابقه بدهی برای شما ایجاد کرده ام، لطفا جزئیات را ملاحظه فرمایید در {{.ReceiptURL}}",
		"it-IT": "Ho creato un debito a tuo nome, controlla i dettagli qui - {{.ReceiptURL}}",
	},
	EMAIL_RECEIPT_BODY_HTML: {
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
		"en-US": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		"fa-IR": "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
		"it-IT": "{{.FromName}} ha creato un debito: {{.ReceiptURL}}",
	},
	INLINE_RECEIPT_TITLE: {
		"ru-RU": "Квитанция: %v",
		"en-US": "Receipt: %v",
		"fa-IR": "رسید: %v",
		"it-IT": "Debito/credito: %v",
	},
	INLINE_RECEIPT_DESCRIPTION: {
		"ru-RU": "Нажмите здесь чтобы отправить квитанцию",
		"en-US": "Click here to send the receipt",
		"fa-IR": "برای ارسال رسید اینجا کلیک کنید.",
		"it-IT": "Clicca qui per inviare un debito/credito",
	},
	INLINE_RECEIPT_CHOOSE_LANGUAGE: {
		"ru-RU": "<b>Выберите язык чтобы посмотреть подробности записи о долге</b> которую создал(а) {{.Creator}}.",
		"en-US": "<b>Please choose language to see details of the debt</b> that has been recorded by {{.Creator}}.",
		"fa-IR": "<b> لطفا برای رویت جزئیات بدهی که توسط </b>  {{.Creator}} ثبت شده است زبان را انتخاب کنید.",
		"it-IT": "<b>Scegli la lingua per vedere i dettagli del debito</b> registrato da {{.Creator}}.",
	},
	INLINE_RECEIPT_MESSAGE: {
		"ru-RU": `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.

{{.SiteLink}} — программа для учёта долгов поможет:

  - Всегда знать кто кому сколько должен

  - Незабыть вовремя отдать или востребовать долг
    <i>(напоминания вам и вашим должникам)</i>`,
		//-------------------------------------------------------
		"en-US": `<b>{{.Creator}} recorded a debt</b> associated with you.

{{.SiteLink}} — an app for debts tracking will help you to:

  - Always know your bottom line

  - Return debts on time
    <i>(reminders to you & your debtors)</i>`,
		//-------------------------------------------------------
		"fa-IR": `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.

{{.SiteLink}} — یک برنامه پیگیری بدهی است که به شما کمک می کند تا:

  - همیشه از سود و زیان خود مطلع باشید.

  - بدهی ها به موقع پرداخت شوند.
    <i>(با ارسال یادآوری به  شما و بدهکاران )</i>`,

		"it-IT": `<b>{{.Creator}} ha registrato un debito</b> associato a te.

{{.SiteLink}} — un app per i debiti che ti consento di:

  - Sapere sempre chi deve soldi a chi

  - Restituire i soldi in tempo
    <i>(lo ricorda a te ed al tuo debitore)</i>`,
	},
	INLINE_INVITE_TITLE: {
		"ru-RU": "Приглашение в %v",
		"en-US": "Invitation to %v",
		"fa-IR": "دعوت به %v",
		"it-IT": "Invito per %v",
	},
	INLINE_INVITE_DESCRIPTION: {
		"ru-RU": "Нажмите здесь для отправки приглашения",
		"en-US": "Click here to send an invite",
		"fa-IR": "برای ارسال یک دعوتنامه اینجا کلیک کنید.",
		"it-IT": "Clicca qui per spedire un invito",
	},
	INLINE_INVITE_MESSAGE: {
		"ru-RU": "%v пригласил вас попробовать %v",
		"en-US": "%v invited you to try %v",
		"fa-IR": "%v شمارا دعوت کرده است به امتحان %v",
		"it-IT": "%v ti ha invitato a provare %v",
	},
	SMS_RECEIPT_YOU_GOT: {
		"ru-RU": "Вы получили %v от %v.",
		"en-US": "You've got %v from %v.",
		"fa-IR": "شما دریافت کرده اید %v از %v.",
		"it-IT": "Hai ricevuto %v da %v",
	},
	SMS_RECEIPT_YOU_GAVE: {
		"ru-RU": "Вы дали %v - взял(а) %v.",
		"en-US": "You've given %v to %v.",
		"fa-IR": "شما پرداخت کرده اید %v به %v.",
		"it-IT": "Hai dato %v a %v",
	},
	SMS_CLICK_TO_CONFIRM_OR_DECLINE: {
		"ru-RU": "Нажмите %v чтобы подтвердить или отвергнуть.",
		"en-US": "Click %v to confirm or decline.",
		"fa-IR": "کلیک کنید %v تا رد خود را تایید نمایید",
		"it-IT": "Clicca %v per confermare o rifiutare",
	},
	HTML_DATE: {
		"ru-RU": "Дата",
		"en-US": "Date",
		"fa-IR": "تاریخ",
		"it-IT": "Data",
	},
	HTML_RECEIPT: {
		"ru-RU": "Квитанция",
		"en-US": "Receipt",
		"fa-IR": "رسید",
		"it-IT": "Scontrino", //To upgrade, not the best translation from Russian

	},
	HTML_AMOUNT: {
		"ru-RU": "Сумма",
		"en-US": "Amount",
		"fa-IR": "مقدار",
		"it-IT": "Totale",
	},
	HTML_FROM: {
		"ru-RU": "Дал",
		"en-US": "From",
		"fa-IR": "از",
		"it-IT": "Da",
	},
	HTML_TO: {
		"ru-RU": "Получил",
		"en-US": "To",
		"fa-IR": "به",
		"it-IT": "Per",
	},
	TELEGRAM_RECEIPT: {
		"ru-RU": "{{.FromName}} создал запись о долге ({{.TransferCurrency}})",
		"en-US": "{{.FromName}} created a debtrecord ({{.TransferCurrency}})",
		"fa-IR": "{{.FromName}} ایجاد یک سابقه بدهی ({{.TransferCurrency}})",
		"it-IT": "{{.FromName}} ha registrato un debito ({{.TransferCurrency}})",
	},
	MESSAGE_TEXT_PLEASE_CHOOSE_FROM_OPTIONS_PROVIDED: {
		"ru-RU": "Пожалуйста выберете из предоставленных опций.",
		"en-US": "Please choose from provided options.",
		"fa-IR": "لطفاً از گزینه های ارائه شده انتخاب نمایید.",
		"it-IT": "Scegli tra le opzioni fornite.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE_OR_COMMENT: {
		"ru-RU": "<b>Хотите добавить заметку или комментарий?</b>\n%v Заметки хранятся для вашего личго пользования.\n%v Комментарий виден всем кому разрешён просмотр этой транзакции.",
		"en-US": "<b>Do you want to add a note or comment?</b>\n%v Memos are private records for yoru own reference.\n%v Comments are available to everyone who has permission to view this transaction.",
		"it-IT": "<b>Vuoi aggiungere una nota o un commento?</b> \n%v I memo sono record privati per il riferimento di yoru.\n%v I commenti sono disponibili a tutti coloro che hanno l'autorizzazione a visualizzare questa transazione.",
		"fa-IR": "<b>آیا میخواهید یادداشت یا شرحی اضافه کنید؟</b>\n%v یادداشت ها نوشته های خصوصی برای مراجعه خود شما هستند.\n%v شرح در دسترس تمام کسانی که مجاز رویت این تراکنش هستند میباشد.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE: {
		"ru-RU": "Напишите заметку:",
		"en-US": "Please write a note:",
		"fa-IR": "لطفاً یک یادداشت بنویسید:",
		"it-IT": "Per favore scrivi un appunto:",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT_ONLY: {
		"ru-RU": `Если хотите добавить комментарий просто отправьте текст.`,
		"en-US": `If you want to add a comment just send a text now.`,

		"fa-IR": `شما می توانید یک شرح اضافه کنید. تنها کافیست یک متن ارسال کنید.`,

		"it-IT": `Se vuoi aggiungere un commento invia del testo ora.`,
	},
	MESSAGE_TEXT_VISIBLE_TO_YOU_AND_COUNTERPARTY: {
		"ru-RU": "виден вам и %v",
		"en-US": "visible to you & %v",
		"fa-IR": "قابل مشاهده برای شما & %v",
		"it-IT": "visibile solo a te e %v",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT: {
		"ru-RU": "Напишите комментарий:",
		"en-US": "Please write the comment:",
		"fa-IR": "لطفاً شرح را ثبت کنید:",
		"it-IT": "Per favore scrivi un commento:",
	},
	MESSAGE_TEXT_TRANSFER_NOTE_ADDED_ASK_FOR_COMMENT: {
		"ru-RU": "Заметка добавлена. Хотите написать комментарий?",
		"en-US": "Memo have been added. Do you want to write a comment?",
		"fa-IR": "یادداشت اضافه شد. آیا میخواهید یک شرح ثبت کنید؟",
		"it-IT": "Promemoria aggiunto. Vuoi scrivere un commento?",
	},
	MESSAGE_TEXT_TRANSFER_COMMENT_ADDED_ASK_FOR_NOTE: {
		"ru-RU": "Комментарий добавлен. Хотите написать заметку?",
		"en-US": "Comment have been added. Do you want to write a note?",
		"fa-IR": "شرح موردنظر شما ثبت شد. آیا می خواهید یک یادداشت بنویسید؟",
		"it-IT": "Commento aggiunto. Vuoi scrivere un appunto?",
	},
	COMMAND_TEXT_NO_COMMENT_OR_NOTE_FOR_TRANSFER: {
		"ru-RU": "Без заметок и комментариев",
		"en-US": "Without notes or comments",
		"fa-IR": "بدون یادداشت یا شرح",
		"it-IT": "Senza appunti o commenti",
	},
	COMMAND_TEXT_NO_COMMENT_FOR_TRANSFER: {
		"ru-RU": "Без комментариев",
		"en-US": "No comments",
		"fa-IR": "بدون شرح",
		"it-IT": "Nessun commento",
	},
	COMMAND_TEXT_NO_NOTE_FOR_TRANSFER: {
		"ru-RU": "Без заметок",
		"en-US": "Without notes",
		"fa-IR": "بدون یادداشت",
		"it-IT": "Senza appunti/note",
	},
	COMMAND_TEXT_ADD_NOTE_TO_TRANSFER: {
		"ru-RU": "Добавить заметку",
		"en-US": "Add a note (private)",
		"fa-IR": "یک یادداشت اضافه کنید (خصوصی)",
		"it-IT": "Aggiungi una nota (privata)",
	},
	COMMAND_TEXT_ADD_COMMENT_TO_TRANSFER: {
		"ru-RU": "Добавить комментарий",
		"en-US": "Add a comment (public)",
		"fa-IR": "یک شرح اضافه کنید (عمومی)",
		"it-IT": "Aggiungi un commento (pubblico)",
	},
	DUE_IN_NOW: {
		"ru-RU": "прямо сейчас",
		"en-US": "now",
		"fa-IR": "حالا",
		"it-IT": "adesso",
	},
	DUE_IN_A_MINUTE: {
		"ru-RU": "через минуту",
		"en-US": "in a minute",
		"fa-IR": "ظرف یک دقیقه",
		"it-IT": "in un minuto",
	},
	DUE_IN_X_MINUTES: {
		"ru-RU": "через %v минут(ы)",
		"en-US": "in %v minutes",
		"fa-IR": "در %v دقیقه",
		"it-IT": "in %v minuti/o",
	},
	DUE_IN_AN_HOUR: {
		"ru-RU": "через час",
		"en-US": "in an hour",
		"fa-IR": "ظرف یک ساعت",
		"it-IT": "in un ora",
	},
	DUE_IN_X_HOURS: {
		"ru-RU": "через %v часа/часов",
		"en-US": "in %v hours",
		"fa-IR": "در %v ساعت",
		"it-IT": "in %v ore/a",
	},
	DUE_IN_X_DAYS: {
		"ru-RU": "через %v дня/дней",
		"en-US": "in %v days",
		"fa-IR": "در %v روز",
		"it-IT": "in %v giorni/o",
	},
	//-------------------------------------------------------------------------------------------------------------------
	WS_ALEX_T: {
		"ru-RU": "Александр Трахимёнок",
		"en-US": "Alexander Trakhimenok",
		"fa-IR": "الکساندر تراخیمِنوک",
		"it-IT": "Alexander Trakhimenok",
	},

	WS_INDEX_TITLE: {
		"ru-RU": "DebtsTracker.io - программа для учёта личных долгов и активов",
		"en-US": "DebtsTracker.io - an IOU app to track your personal debts & assets",
		"es-ES": "DebtsTracker.io - una aplicación para el seguimiento de sus deudas personales",
		"fa-IR": "DebtsTracker.io - برنامه ای برای ردیابی بدهی ها و دارایی های شما",
		"pl-PL": "DebtsTracker.io - aplikacja do śledzenia osobistych długów",
		"pt-PT": "DebtsTracker.io - um aplicativo para controlar suas dívidas pessoais",
		"de-DE": "DebtsTracker.io - eine App, um Ihre persönlichen Schulden zu verfolgen",
		"fr-FR": "DebtsTracker.io - une application pour suivre vos dettes personnelles",
		"it-IT": "DebtsTracker.io - un app per monitorare i tuoi debiti personali",
		"ko-KO": "DebtsTracker.io 은 - 앱이 사용자의 개인 채무를 추적",
		"ja-JP": "DebtsTracker.io は - アプリはあなたの個人的な借金を追跡します",
		"zh-CN": "DebtsTracker.io - 一个应用程序来跟踪你的个人债务",
	},
	WS_LIVE_DEMO: {
		"ru-RU": "Демо версия online",
		"en-US": "Live demo",
		"es-ES": "Demo en vivo",
		"fa-IR": "نسخه ی نمایشی زنده",
		"pl-PL": "Demo na żywo",
		"pt-PT": "Demonstração ao vivo",
		"de-DE": "Echtzeit Vorführung",
		"fr-FR": "Démo en direct",
		"it-IT": "Demo online",
		"ko-KO": "실시간 데모",
		"ja-JP": "ライブデモ",
		"zh-CN": "现场演示",
	},
	WS_INDEX_TG_BOT_H2: {
		"ru-RU": "Бот для Telegram",
		"en-US": "Chat bot for Telegram messenger",
		"es-ES": "Chat bot para Telegrama mensajero",
		"fa-IR": "ربات چت برای پیام رسان تلگرام",
		"pl-PL": "Chat bot do telegramu posłańca",
		"pt-PT": "bot de bate-papo para Telegram messenger",
		"de-DE": "Chat-Bot für Telegrammbote",
		"fr-FR": "bot Chat for Telegram messenger",
		"it-IT": "Bot Chat per Telegram",
		"ko-KO": "전보 메신저 채팅 봇",
		"ja-JP": "電報メッセンジャーのためのチャットボット",
		"zh-CN": "聊天机器人的电报使者",
	},
	WS_INDEX_TG_BOT_OPEN: {
		"ru-RU": "Открыть в Телеграмме &#x1F680;",
		"en-US": "Open in Telegram &#x1F680;",
		"es-ES": "Abrir en Telegrama &#x1F680;",
		"fa-IR": "بازکردن در تلگرام &#x1F680;",
		"pl-PL": "Otwórz w telegramu &#x1F680;",
		"pt-PT": "Open in Telegram &#x1F680;",
		"de-DE": "Open in Telegramm &#x1F680;",
		"fr-FR": "Open in Telegram &#x1F680;",
		"it-IT": "Apri su Telegram &#x1F680;",
		"ko-KO": "전보 &#x1F680; 에서 열기;",
		"ja-JP": "電報 &#x1F680; で開きます。",
		"zh-CN": "打开在电报 &#x1F680;",
	},

	WS_INDEX_TG_BOT_P: {
		"ru-RU": "В настоящий момент наша программа доступна в мессенджере <a href='https://telegram.org/'>Телеграм</a>.",
		"en-US": "At the moment our program is available just on <a href='https://telegram.org/'><b>Telegram</b> messenger</a>",
		"es-ES": "Por el momento nuestro programa está disponible sólo en <a href='https://telegram.org/'> <b> Telegrama </b> mensajero </a>",
		"fa-IR": "درحال حاضر برنامه ما فقط در دسترس است در <a href='https://telegram.org/'>تلگرام</a>",
		"pl-PL": "W tej chwili nasz program jest dostępny tylko na <a href='https://telegram.org/'> <b> Telegram </b> messenger </a>",
		"pt-PT": "No momento em que o nosso programa está disponível apenas na <a href='https://telegram.org/'> <b> Telegram </b> messenger </a>",
		"de-DE": "Im Moment ist unser Programm nur auf verfügbar <a href='https://telegram.org/'> <b> Telegramm </b> Bote </a>",
		"fr-FR": "Au moment de notre programme est disponible seulement sur <a href='https://telegram.org/'> <b> Telegram </b> messager </a>",
		"it-IT": "Al momento il nostro programma è disponibile solo su <a href='https://telegram.org/'> <b> Telegram </b></a>",
		"ko-KO": "지금이 순간 우리의 프로그램은 단지에 <a href='https://telegram.org/'>의 <b> 전보 </b>을 메신저 </a>를 볼 수 있습니다",
		"ja-JP": "現時点では私たちのプログラムは、ちょうど上の<a href='https://telegram.org/'><B>電報</b>のメッセンジャー</a>で提供されています",
		"zh-CN": "目前我们的计划是只提供在<a href='https://telegram.org/'><B>电报</b>的使者</A>",
	},
	WS_MOTTO: {
		"ru-RU": "Платежи по долгам целиком и вовремя!",
		"en-US": "Know your bottom line & never miss a debt payment!",
		"es-ES": "Conozca a su línea de fondo y nunca se pierda un pago de la deuda!",
		"fa-IR": "از سود و زیان خود مطلع باشید و هرگز پرداخت بدهی ای را از قلم نیندازید",
		"pl-PL": "Znaj swoją równowagę i nigdy nie przegapisz zapłatę długu!",
		"pt-PT": "Conheça o seu equilíbrio e nunca perca um pagamento da dívida!",
		"de-DE": "Ihr Kontostand wissen und nie eine Schuld Zahlung verpassen!",
		"fr-FR": "Apprenez à connaître votre solde et ne jamais manquer un paiement de la dette!",
		"it-IT": "Tieni sott'occhio il tuo bilancio e non dimenticarti mai di un debito!",
		"ko-KO": "균형을 알고 및 채무 지불을 놓칠 수 없어!",
		"ja-JP": "あなたのバランスを知っている＆債務の支払いを見逃すことはありません！",
		"zh-CN": "了解天平＆不会错过任何一个债务付款！",
	},
	WS_SHORT_DESC: {
		"ru-RU": "DebtsTracker.io - мобильное приложение и сервис напоминаний для учёта и своевременной выплаты долгов. Отсылает автоматические уведомления вашим должникам по email и SMS.",
		"en-US": "DebtsTracker.io is a mobile IOU app & a reminder service that helps to track your debts, credits & assets. Sends automated email & SMS reminders to your debtors.",
		"es-ES": "DebtsTracker.io es un servicio de aplicaciones móviles y recordatorio de que ayuda a realizar un seguimiento de sus deudas y créditos. Envía notificaciones por correo electrónico y SMS automatizados a sus deudores.",
		"fa-IR": "DebtsTracker.io یک برنامه موبایل و سرویس یادآور می باشد که به شما کمک می کند تا بدهی ها و اعتبارات خود را ردیابی نمایید. همچنین ایمیل و پیام کوتاه یادآوری به بدهکاران ارسال می کند.",
		"pl-PL": "DebtsTracker.io to aplikacje mobilne i przypomnienia usługa, która pozwala na śledzenie swoich długów i kredytów. Wysyła automatycznych powiadomień e-mail i SMS do swoich dłużników.",
		"pt-PT": "DebtsTracker.io é um serviço de aplicativos móveis e lembrete de que ajuda a controlar seus débitos e créditos. Envia e-mail e SMS notificações automáticas aos seus devedores.",
		"de-DE": "DebtsTracker.io ist eine mobile Apps und Erinnerungs-Service, die Ihre Schulden und Kredite zu verfolgen hilft. Sendet automatisierte E-Mail und SMS-Benachrichtigungen an Ihre Schuldner.",
		"fr-FR": "DebtsTracker.io est une des applications mobiles et rappel service qui permet de suivre vos dettes et crédits. Envoie automatisés email & SMS reminders à vos débiteurs.",
		"it-IT": "DebtsTracker.io è un servizio di applicazioni mobili che ricordare e aiuta a monitorare i debiti e crediti. Invia notifiche e-mail e SMS automatici ai i vostri debitori.",
		"ko-KO": "DebtsTracker.io 은 채무 및 크레딧을 추적하는 데 도움이 모바일 앱 및 알림 서비스입니다. 당신의 채무자에 자동화 된 이메일 및 SMS 알림을 보냅니다.",
		"ja-JP": "DebtsTracker.io は、あなたの借金＆クレジットを追跡するのに役立ちますモバイルアプリ＆リマインダーサービスです。あなたの債務者に自動メール＆SMS通知を送信します。",
		"zh-CN": "DebtsTracker.io 是一个移动应用和提醒服务，帮助跟踪你的债务和信用。发送自动电子邮件和短信通知到您的债务人。",
	},

	WS_HELP_US_TITLE: {
		"en-US": "How you can help to DebtsTracker.io project",
		"ru-RU": "Как вы можете помочь проекту DebtsTracker.io",
		"it-IT": "Come potete aiutare il progetto DebtsTracker.io", // TODO(IT): Google translated
		"fa-IR": "چگونه می توانید به پروژه  DebtsTracker.io کمک کنید.",
	},
	WS_ADS_TITLE: {
		"en-US": "Ads @ DebtsTracker.IO",
		"ru-RU": "Реклама на DebtsTracker.IO",
		"fa-IR": "تبلیغات @ DebtsTracker.IO",
		"it-IT": "Pubblicita' @ DebtsTracker.IO",
	},
	WS_ADS_CONTENT: {
		"en-US": `To place ads in our app please send us an email to <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"ru-RU": `Чтобы разместить рекламу в нашем приложении пишите нам на <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"fa-IR": `برای قراردادن تبلیغات در برنامه ما، درخواست خود را به این آدرس ایمیل فرمایید <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"it-IT": `Per inserire della pubblicita nella nostra app inviaci un email a <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
	},
	FB_MAKE_RECORD_HEADER: {
		"en-US": "Record debt",
		"ru-RU": "Записать долг",
		"fa-IR": "ثبت بدهی",
		"it-IT": "Registra il debito",
	},
	FB_MAKE_RECORD_HEADLINE: {
		"en-US": "Scroll left to see other options.",
		"ru-RU": "Пролистайте карточки влево чтобы увидеть другие опции.",
		"fa-IR": "برای دیدن سایر گزینه ها به سمت چپ اسکرول نمایید.",
		"it-IT": "Scrolla a sinistra per vedere altre opzioni",
	},
	FB_HOW_ARE_THINGS_HEADER: {
		"en-US": "How are you doing?",
		"ru-RU": "Как идут дела?",
		"fa-IR": "حال شما چطوره؟",
		"it-IT": "Come te la passi?",
	},
}
