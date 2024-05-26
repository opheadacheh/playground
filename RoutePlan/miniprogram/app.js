// app.js
App({
  onLaunch() {
    // 展示本地存储能力
    const logs = wx.getStorageSync('logs') || []
    logs.unshift(Date.now())
    wx.setStorageSync('logs', logs)

    // 登录
    wx.login({
      success: function(res) {
        // 发送 res.code 到后台换取 openId, sessionKey, unionId
      }
    })
  },
  globalData: {
    searchPlaceHolders: [
        "火锅",
        "自助餐",
        "烧烤",
        "KTV",
        "按摩",
        "卡丁车",
        "室内攀岩",
        "射箭",
        "游乐场",
        "爬山",
        "湘菜",
        "日料",
        "韩国烤肉",
    ],
    tagMap: new Map([
        ["LEAST_EXPENSIVE", "最实惠"],
        ["HIGHEST_RATING", "评分最高"],
    ])
  }
})
