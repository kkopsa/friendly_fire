package com.example.kevin.myapplication;

import android.annotation.TargetApi;
import android.content.BroadcastReceiver;
import android.content.Context;
import android.content.Intent;
import android.os.Build;
import android.provider.Telephony;
import android.telephony.SmsMessage;
import android.util.Log;
import android.widget.Toast;

import static android.widget.Toast.LENGTH_SHORT;
import static android.widget.Toast.makeText;

/**
 * Created by kevin on 6/26/15.
 */
public class SMSBroadcastReceiver extends BroadcastReceiver {

    private String mMessageBody; // stores message body
    private String mGPScoords = "Hello World!"; // stores gps coordinates

    @TargetApi(Build.VERSION_CODES.KITKAT)
    @Override
    public void onReceive(Context context, Intent intent) {
        if (Telephony.Sms.Intents.SMS_RECEIVED_ACTION.equals(intent.getAction())) {
            for (SmsMessage smsMessage : Telephony.Sms.Intents.getMessagesFromIntent(intent)) {
                mMessageBody = smsMessage.getMessageBody();
                Log.i("SMSmessage", mMessageBody);
            }
            Toast message = makeText(context, "Orders recieved", LENGTH_SHORT);
            message.show();
            parseSMSMessage();

        }
    }

    //message parser for GPS coordinates
    public void parseSMSMessage(){
        boolean coordRecieved = false; // have we recieved the GPS coordinates?
        mGPScoords = ""; // clear String
        for (int i = 0; !coordRecieved; i++) {
            // if you find a number
            if (mMessageBody.charAt(i) >= '0' &&
                    mMessageBody.charAt(i) <= '9') {
                // add characters to the string until a space is found
                for (int j = i; mMessageBody.charAt(j) != ' '; j++) {
                    mGPScoords += mMessageBody.charAt(j);
                }
                coordRecieved = true; // we got what we came for
                Log.e("GPSCoordinates", mGPScoords); // DEBUG: Check it!
            }
        }
    }
}