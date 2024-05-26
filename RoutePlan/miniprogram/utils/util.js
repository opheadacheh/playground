var app = getApp()

function getRandomPlaceHolder() {
    const length = app.globalData.searchPlaceHolders.length
    return app.globalData.searchPlaceHolders[Math.floor(Math.random() * length)]
}

module.exports = {
    getRandomPlaceHolder: getRandomPlaceHolder,
}
