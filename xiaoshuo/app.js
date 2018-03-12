//app.js 
var https = require('utils/https.js')
var url = "http://192.168.0.200:8888/";


function Login(code, userInfo) {
  var that = this; 
  var nickName = userInfo.nickName;
  var avatarUrl = userInfo.avatarUrl;
  var gender = userInfo.gender;
  var language = userInfo.language;
  var loginUrl =url+"user/login?code="+code;
  https.get(loginUrl, function (res) {
    var token = res.token;
    wx.setStorageSync('token', token);
  },function(){
     //wx.setStorageSync('token', "defalutToken")
     // 在这里你要考虑到用户登录失败的情况 
  },false)
} 


App({
  globalData: {
    domain: url,
    detailUrl:url+"novels/get",
    novelsUrl: url+"novels",
    recommendUrl: url+"novels/recommend",
    chaptersUrl: url+"novels/getChapters",
    chaptersContentUrl: url+"novels/getChaptersContent",
    categoryUrl: url+"novels/getCategory",
    searchUrl: url+"novels/search",
    feedbackUrl: url+"feedback/send",
    novelSerialNumbersUrl: url+"novels/getNovelSerialNumbers",
    searchKeywords: url+"novels/getSearchKeywords",
    token:null,
    userInfo: null
  },
   onLaunch: function () {
    wx.clearStorage();
    this.getUserInfo();
    this.removeCache();
  },
  removeCache:function(){
    var value = wx.getStorageSync('cateListTime')
    var nowTime = new Date().getTime();
    if(value){
      if (value < nowTime) {
        wx.removeStorage("cateList");
      }
    }
    value = wx.getStorageSync('searchKeysTime')
    if (value) {
      if (value < nowTime) {
        wx.removeStorage("searchKeysTime");
      }
    }
  },
  WechatLogin:function(){
    var that = this
    wx.login({
      success: function (res) {
        //登录成功 
        if (res.code) {
          // 这里是用户的授权信息每次都不一样 
          var code = res.code;
          wx.getUserInfo({
            success: function (info) {
              wx.setStorageSync('userInfo', info.userInfo);
              that.globalData.userInfo = info.userInfo;
              typeof cb == "function" && cb(that.globalData.userInfo)
              // 请求自己的服务器
              Login(code, info.userInfo);
            },fail:function(e){
              console.log(e);
            }
          })
        }
      }, fail: function (e) {
         console.log(e);
      }
    })
  },
  getUserInfo: function (cb) {
    var that = this
    if (this.globalData.userInfo) {
      console.log(cb);
      typeof cb == "function" && cb(this.globalData.userInfo)
    } else {
        wx.checkSession({
        success: function () {
          var userInfo=wx.getStorageSync('userInfo');
          if(userInfo){
            that.globalData.userInfo = userInfo;
          }else{
            that.WechatLogin();
          }
          console.log("success login");
        },fail(){
          that.WechatLogin();
        }});
    } 
   }
})