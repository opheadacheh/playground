// pages/route/route.js
Page({

    /**
     * 页面的初始数据
     */
    data: {
        businesses: [],
    },

    /**
     * 生命周期函数--监听页面加载
     */
    onLoad(options) {
        var route_data = wx.getStorageSync('route')
        console.log(route_data)
        this.setData({
            businesses: route_data.businesses,
        })
    },

    onViewInfoTapped: function(e) {
        console.log(e)
        var currentIndex = e.currentTarget.dataset.index
        wx.openLocation({
          latitude: this.data.businesses[currentIndex].latitude,
          longitude: this.data.businesses[currentIndex].longitude,
        })
    }
})