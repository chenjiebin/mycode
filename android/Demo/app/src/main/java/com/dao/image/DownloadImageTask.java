package com.dao.image;

import android.content.Context;
import android.os.AsyncTask;
import android.graphics.Bitmap;
import android.util.Log;

import org.apache.http.client.HttpClient;
import org.apache.http.impl.client.DefaultHttpClient;

/**
 * Created by chenjiebin on 1/16/15.
 */
public class DownloadImageTask extends AsyncTask<String, Integer, Bitmap> {
    private Context mContext;

    DownloadImageTask(Context context) {
        mContext = context;
    }

    @Override
    protected void onPreExecute() {
        super.onPreExecute();
    }

    @Override
    protected Bitmap doInBackground(String... urls) {
        Log.v("doInBackground", "doing download of image");
        return null;
    }

    private Bitmap downloadImage(String... urls) {
        HttpClient client = new DefaultHttpClient();


        return null;
    }
}
