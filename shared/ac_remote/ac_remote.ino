#include <ESP8266WiFi.h>
#include <WiFiClient.h>
#include <ESP8266WebServer.h>
#include <ESP8266mDNS.h>

#include <Arduino.h>
#include <IRremoteESP8266.h>
#include <IRsend.h>
#include <IRrecv.h>
#include <ir_Samsung.h>
#include <IRac.h>
#include <IRtext.h>
#include <IRutils.h>


#ifndef STASSID
#define STASSID ""
#define STAPSK  ""
#endif


/*
 * IR Module setup
 */

uint16_t ac_off_data[349] = {600, 17772,  2984, 8942,  536, 478,  490, 1476,  512, 504,  488, 504,  490, 504,  490, 502,  492, 500,  492, 502,  492, 508,  484, 1500,  512, 504,  512, 482,  488, 1476,  508, 1478,  512, 504,  512, 1448,  514, 1498,  514, 1450,  536, 1472,  516, 1470,  518, 504,  486, 508,  464, 528,  488, 506,  486, 506,  488, 506,  486, 506,  484, 508,  512, 482,  512, 480,  490, 502,  512, 480,  514, 478,  490, 502,  514, 480,  490, 500,  516, 478,  516, 478,  492, 508,  510, 504,  488, 504,  486, 508,  512, 482,  488, 504,  488, 504,  490, 504,  516, 478,  516, 476,  492, 500,  492, 502,  516, 476,  514, 478,  518, 476,  520, 474,  518, 1468,  518, 1474,  512, 2950,  3012, 8940,  538, 1474,  514, 478,  514, 478,  512, 480,  520, 472,  520, 472,  520, 474,  520, 474,  518, 474,  520, 1468,  518, 476,  520, 472,  546, 1440,  550, 472,  522, 1464,  514, 1472,  516, 1470,  514, 1474,  546, 1440,  546, 1440,  550, 444,  550, 442,  552, 442,  548, 444,  552, 442,  550, 442,  552, 442,  552, 468,  524, 470,  524, 468,  512, 480,  514, 478,  514, 478,  544, 450,  548, 446,  580, 412,  580, 414,  580, 412,  558, 436,  554, 438,  556, 436,  558, 434,  558, 434,  560, 434,  558, 436,  558, 460,  532, 460,  492, 500,  508, 486,  538, 454,  538, 454,  552, 440,  580, 414,  582, 410,  558, 434,  560, 432,  584, 2884,  3078, 8846,  610, 1402,  586, 408,  584, 410,  584, 410,  582, 436,  556, 436,  556, 436,  476, 516,  506, 486,  540, 1444,  578, 416,  560, 434,  560, 1424,  586, 1402,  584, 1402,  562, 1424,  562, 432,  560, 1426,  562, 1452,  534, 1452,  508, 1480,  538, 1448,  556, 1432,  586, 1400,  586, 1400,  584, 410,  584, 410,  582, 410,  582, 1402,  586, 1400,  586, 1430,  558, 436,  526, 466,  504, 490,  538, 454,  548, 446,  552, 440,  560, 434,  560, 434,  558, 1406,  582, 1406,  576, 438,  556, 1408,  604, 1402,  584, 414,  580, 438,  556, 438,  550, 442,  528, 464,  478, 514,  510, 482,  540, 452,  554, 440,  556, 436,  558, 1430,  560, 1406,  572};  // SAMSUNG_AC
uint8_t ac_off_state[21] = {0x02, 0xB2, 0x0F, 0x00, 0x00, 0x00, 0xC0, 0x01, 0xD2, 0x0F, 0x00, 0x00, 0x00, 0x00, 0x01, 0xF2, 0xFE, 0x71, 0x80, 0x0D, 0xC0};

uint16_t ac_on_data[349] = {620, 17822,  3008, 8918,  536, 480,  512, 1470,  516, 480,  512, 480,  512, 480,  512, 478,  514, 480,  512, 480,  512, 508,  484, 1500,  512, 484,  458, 534,  482, 1502,  512, 482,  512, 482,  512, 1472,  516, 1472,  514, 1452,  536, 1472,  514, 1472,  514, 480,  512, 508,  486, 508,  512, 480,  512, 482,  458, 534,  484, 510,  484, 508,  512, 482,  512, 480,  512, 480,  512, 480,  512, 480,  512, 482,  512, 480,  514, 480,  512, 480,  514, 480,  512, 480,  514, 506,  486, 508,  512, 480,  510, 482,  458, 534,  482, 512,  484, 510,  510, 480,  512, 482,  512, 480,  512, 482,  510, 482,  512, 480,  514, 1474,  512, 1474,  514, 1472,  514, 1500,  512, 2930,  3006, 8946,  482, 1526,  486, 510,  486, 508,  510, 482,  512, 480,  512, 482,  512, 480,  512, 482,  512, 480,  514, 1472,  514, 480,  512, 480,  512, 1500,  488, 508,  458, 1526,  514, 1472,  486, 1502,  510, 1476,  514, 1452,  536, 1472,  514, 482,  512, 480,  514, 480,  512, 480,  512, 480,  512, 480,  512, 508,  486, 508,  488, 504,  510, 482,  510, 482,  482, 510,  482, 510,  510, 482,  512, 482,  512, 480,  512, 480,  512, 480,  512, 482,  512, 480,  512, 480,  514, 480,  512, 480,  512, 480,  538, 484,  484, 508,  460, 534,  512, 482,  458, 534,  482, 512,  484, 508,  512, 482,  512, 480,  512, 480,  512, 480,  514, 480,  512, 2954,  3008, 8918,  534, 1474,  512, 482,  510, 480,  540, 454,  512, 508,  484, 508,  510, 482,  458, 534,  458, 534,  484, 1500,  540, 454,  514, 480,  538, 1446,  514, 482,  512, 1474,  538, 1450,  514, 480,  512, 1474,  512, 1498,  488, 1498,  512, 1474,  462, 1526,  488, 1498,  516, 1472,  516, 1472,  514, 480,  514, 480,  514, 480,  538, 1448,  512, 1474,  514, 1498,  492, 504,  458, 534,  510, 484,  510, 482,  482, 510,  486, 506,  512, 482,  512, 480,  514, 1468,  518, 1470,  516, 480,  538, 1448,  514, 1472,  512, 482,  512, 506,  486, 506,  460, 532,  508, 484,  510, 482,  482, 510,  486, 508,  514, 1470,  518, 1468,  518, 1470,  516, 1472,  514};  // SAMSUNG_AC
uint8_t ac_on_state[21] = {0x02, 0x92, 0x0F, 0x00, 0x00, 0x00, 0xF0, 0x01, 0xD2, 0x0F, 0x00, 0x00, 0x00, 0x00, 0x01, 0xD2, 0xFE, 0x71, 0x80, 0x0D, 0xF0};

const uint16_t kIrLed = 4;  // Transmiter GPIO PIN: 4 (D2 on NODEMCU).
const uint16_t kRecvPin = 14; // Receiver GPIO PIN: 16 (D5 on NODEMCU).
const uint16_t kCaptureBufferSize = 1024;
const uint8_t kTimeout = 50;
const uint8_t kTolerancePercentage = kTolerance;  // kTolerance is normally 25%
const uint16_t kMinUnknownSize = 12;


IRsend irsend(kIrLed);
IRrecv irrecv(kRecvPin, kCaptureBufferSize, kTimeout, true);
decode_results results;  // Somewhere to store the results

ESP8266WebServer server(80);

void setup() {
  // put your setup code here, to run once:
  Serial.begin(115200);
  WiFi.mode(WIFI_STA);
  WiFi.begin(STASSID, STAPSK);
  
  Serial.println();
  Serial.print("connecting to WiFi");
  Serial.print(STASSID);
  Serial.println();
  
  while (WiFi.status() != WL_CONNECTED) {
     delay(500);
     Serial.print(".");
  }

  Serial.println("");
  Serial.print("Connected! IP: ");
  Serial.println(WiFi.localIP());
  
  if (MDNS.begin("esp8266")) {
    Serial.println("MDNS responder started");
  }

  server.on("/ac/on", handleLivingRoomACON);
  server.on("/ac/off", handleLivingRoomACOFF);
  
  irsend.begin();
  server.begin();
  irrecv.setTolerance(kTolerancePercentage);  // Override the default tolerance.
  irrecv.enableIRIn();  // Start the receiver

}

void handleLivingRoomACON() {
    Serial.println("Turn on the A/C ...");
    irsend.sendRaw(ac_on_data, 349, 38);
    server.send(200, "text/plain", "LIVING_ROOM_AC_ON");
}
void handleLivingRoomACOFF() {
    Serial.println("Turn off the A/C ...");
    irsend.sendRaw(ac_off_data, 349, 38);
    server.send(200, "text/plain", "LIVING_ROOM_AC_OFF");  
}

void decodeIR() {
  if (irrecv.decode(&results)) {
    uint32_t now = millis();
    Serial.print(resultToHumanReadableBasic(&results));
    Serial.println();    // Blank line between entries
  }
}

void loop() {
  server.handleClient();
  MDNS.update();
  decodeIR();
}
