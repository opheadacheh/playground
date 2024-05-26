// pages/routes/routes.js
Page({

    /**
     * 页面的初始数据
     */
    data: {
        routes: [],
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        var routes_data = wx.getStorageSync('routes_data')
        this.setData({
            routes: routes_data,
        })
    },

    onRouteTapped: function(e) {
        var currentIndex = e.currentTarget.dataset.index
        wx.setStorageSync('route', this.data.routes[currentIndex])
        wx.navigateTo({
          url: '/pages/route/route',
        })
    }
})