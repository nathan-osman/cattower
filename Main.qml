import QtQuick
import QtQuick.Controls
import QtQuick.Layouts

ApplicationWindow {
    id: appWindow

    visible: true
    visibility: "FullScreen"

    readonly property int margin: 12

    RowLayout {
        anchors.fill: parent
        anchors.margins: appWindow.margin

        Button {
            icon.source: "content/images/home.svg"
        }

        Button {
            icon.source: "content/images/wifi.svg"
        }

        Button {
            icon.source: "content/images/settings.svg"
        }
    }
}
