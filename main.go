package main

import (
	"log"
	"github.com/dragonmaster101/bob/classifier"
	audio "github.com/dragonmaster101/go_audio/audio"
	chat "github.com/dragonmaster101/go_chat/chat"
	sr "github.com/dragonmaster101/go_speechrecognition/speechrecognition"
	tts "github.com/dragonmaster101/go_texttospeech/watson_tts"
)

const CHAT_TOKEN = "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";
const CHAT_URL = "https://api-inference.huggingface.co/models/facebook/blenderbot-400M-distill"

const TTS_TOKEN = "lDCDgAM735uy2hNDlLgf4hmYnZMMxq9C-8noYJbGjHZQ";
const TTS_URL = "https://api.us-south.text-to-speech.watson.cloud.ibm.com/instances/e83db35d-6925-4f4e-afc6-a6c41ebedff6";

const SR_TOKEN = "djFD3P479m7_iCNgr0pEdBT78GM_Jzg-b0NUnEYwOPgP";
const SR_URL = "https://api.us-south.speech-to-text.watson.cloud.ibm.com/instances/085b5a01-51de-4edb-9c3f-7f118a127863"


func main(){
	convo := chat.Conversation{};
	convo.Init(chat.BasicConversationOption(CHAT_TOKEN , CHAT_URL));
	convo.CreateLog("test");

	go classifier.DisplayDetection(0 , "classifier.xml");

	recognizer := sr.RecognizerIBM(SR_TOKEN , SR_URL);
	ttsService := tts.CreateService(TTS_TOKEN , TTS_URL);

	for {
		audio.Record("test.aiff" , 4);
		log.Println("Recording complete");
		transcription , err := recognizer.TranscribeFile("test.aiff");
		if err != nil {
			log.Fatal(err);
		}
		log.Println(transcription);
		reply := convo.Query(transcription);	
		tts.Synthesize(ttsService , reply , "reply.mp3");
		audio.PlayMp3("reply.mp3");
	}
}