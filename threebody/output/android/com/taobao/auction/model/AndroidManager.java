package com.taobao.auction.model;

import taobao.auction.base.network.HttpHelper;
import taobao.auction.base.network.HttpResponse;

/**
 * 这是一个自动生成的model文件，不要对该文件做任何修改
 * 作者：shanks
 * 时间：2015-08-31 23:56:35
 */
public class AndroidManager {

    public HttpResponse<Android> queryAndroid(long albumId,long msgId) {
        AndroidRequest request = new AndroidRequest();
        
        request.albumId = albumId;
        
        request.msgId = msgId;
        
        HttpResponse<Android> httpResponse = HttpHelper.request(request, Android.class);
        return httpResponse;
    }

}
