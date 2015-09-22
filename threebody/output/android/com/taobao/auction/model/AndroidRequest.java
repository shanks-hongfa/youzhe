package com.taobao.auction.model;

import mtopsdk.mtop.domain.IMTOPDataObject;

/**
 * 这是一个自动生成的model文件，不要对该文件做任何修改
 * 作者：shanks
 * 时间：2015-08-31 23:56:35
 */
public class AndroidRequest implements IMTOPDataObject {
    /**
     * API的名称
     * (Required)
     */
    public String API_NAME = "mtop.taobao.paimai.tbp2.getAllForLive";
    /**
     * API的版本号
     * (Required)
     */
    public String VERSION = "1.0";

    /**
     * API的签名方式
     * (Required)
     * 需要登陆就设置为true 不需要登陆就设置为false
     */
    public boolean NEED_ECODE = true;

    /**
     * 淘宝无线用户会话ID 是否需要带sid
     * (Required)
     * 当NEED_ECODE=true时，NEED_SESSION不需要设置。
     * <p/>
     * 在我看来 这个参数不需要设置，非登录或者登陆，所以一般只需要设置NEED_ECODE就行
     */
    private boolean NEED_SESSION=true;


     
    public long albumId;
    
    public long msgId;
    
}
