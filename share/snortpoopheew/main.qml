import QtQuick 2.0
import Ubuntu.Components 0.1
import QtQuick.Dialogs 1.1

/*!
    \brief MainView with a Label and Button elements.
*/
import QtQuick.Controls 1.1
import Ubuntu.PerformanceMetrics 1.0

MainView {
    id: mainPoopView
    // objectName for functional testing purposes (autopilot-qt5)
    objectName: "mainView"

    // Note! applicationName needs to match the "name" field of the click manifest
    applicationName: "com.ubuntu.developer..goqmltest"

    /*
     This property enables the application to change orientation
     when the device is rotated. The default is false.
    */
    //automaticOrientation: true

    width: units.gu(100)
    height: 400


    FileDialog {
        id: destinationFileDialog
        selectExisting: false
        selectFolder: true
        onAccepted: {
           console.log("destination chosen: " + destinationFileDialog.fileUrl)
           cfg.destinationUrl = destinationFileDialog.fileUrl
        }
    }

    FileDialog {
        id: sourceFileDialog
        selectExisting: true
        selectFolder: true
        onAccepted: {
           console.log("source chosen: " + sourceFileDialog.fileUrl)
           cfg.sourceUrl = sourceFileDialog.fileUrl
        }
    }

    Page {
        id: page1
        y: 268
        height: 368
        anchors.verticalCenterOffset: 0
        anchors.rightMargin: 0
        anchors.leftMargin: 0
        anchors.topMargin: 0
        anchors.verticalCenter: parent.verticalCenter
        anchors.top: parent.top
        anchors.right: parent.right
        anchors.left: parent.left
        title: i18n.tr("SnortPoopPhew")
        x: 200

        Image {
            id: image1
            x: 225
            y: 8
            width: 35
            height: 44
            opacity: 1
            antialiasing: true
            fillMode: Image.PreserveAspectFit
            source: "linse.png"
        }


        Text {
            id: text1
            x: 8
            y: 94
            width: 784
            height: 34
            text: qsTr("Offload your  material and be sure about the files copied...")
            verticalAlignment: Text.AlignVCenter
            horizontalAlignment: Text.AlignHCenter
            font.pixelSize: 12
        }

        Grid {
            id: grid1
            x: 8
            width: 386
            anchors.top: parent.top
            anchors.topMargin: 144
            anchors.bottom: parent.bottom
            anchors.bottomMargin: 342
            spacing: 10
            columns: 0
            rows: 0

            Column {
                x: 8
                y: -128
                visible: true
                anchors.right: parent.right
                anchors.rightMargin: 16
                anchors.left: parent.left
                anchors.leftMargin: 16
                anchors.bottom: parent.bottom
                anchors.bottomMargin: 16
                anchors.top: parent.top
                anchors.topMargin: 16
                spacing: units.gu(1)
                anchors {
                    margins: units.gu(2)
                }


                Button {
                    objectName: "button"
                    width: parent.width
                    text: i18n.tr("Source Directory")
                    onClicked: {
                        sourceFileDialog.setTitle("chosse source directory...")
                        sourceFileDialog.open()
                   }
                }
                
                Label {
                    id: label
                    objectName: "label"
                    text: cfg.sourceUrl
                }
            }
        }


        Grid {
            id: grid2
            x: 406
            width: 386
            anchors.top: parent.top
            anchors.topMargin: 144
            anchors.bottomMargin: 342
            anchors.bottom: parent.bottom
            Column {
                anchors.margins: units.gu(2)
                anchors.rightMargin: 16
                anchors.left: parent.left
                anchors.bottomMargin: 16
                anchors.topMargin: 16
                visible: true
                anchors.right: parent.right
                anchors.leftMargin: 16
                anchors.bottom: parent.bottom
                anchors.top: parent.top
                

                Button {
                    objectName: "button"
                    width: parent.width
                    text: i18n.tr("Destination Directory")
                    onClicked: {
                        destinationFileDialog.setTitle("chosse destination directory...")
                        destinationFileDialog.open()
                    }
                }
                
                Label {
                    id: label1
                    text: cfg.destinationUrl
                    objectName: "label"
                }
                spacing: units.gu(1)
            }
            columns: 0
            spacing: 10
            rows: 0
        }


        GroupBox {
            id: groupBox1
            x: 27
            y: 221
            width: 367
            height: 64
            flat: true
            checked: false
            title: qsTr("Select hash...")

            ExclusiveGroup { id: hashGroup }

            RadioButton {
                id: radioButton1
                x: 100
                y: 0
                text: "MD5"
                checked: cfg.hasMd5() ? true : false
                exclusiveGroup: hashGroup
                onClicked: {
                  console.log("md5 checked")
                  cfg.hash = "MD5"
                }
            }

            RadioButton {
                id: radioButton2
                text: "CRC32"
                checked: cfg.hasCrc() ? true : false
                exclusiveGroup: hashGroup
                onClicked: {
                  console.log("crc32 checked")
                  cfg.hash = "CRC32"
                }
            }
        }

    }


    StatusBar {
        id: statusBar1
        x: 8
        y: 365
        width: 784
        height: 27
        activeFocusOnTab: true

        Text {
            id: statusText
            text: ctrl.statusText
            verticalAlignment: Text.AlignVCenter
        }
    }

    Button {
        id: startCopyBtn
        x: 584
        y: 310
        width: 191
        height: 27
        text: qsTr("Process...")
        onClicked: {
          ctrl.startCopy(cfg)
          startCopyBtn.enabled = false
          cancelCopyBtn.enabled = true
        }
    }

    Button {
      id: cancelCopyBtn
      x: 381
      y: 311
      width: 191
      height: 27
      text: qsTr("Cancel")
      enabled: false
      onClicked: {
        console.log("cancel klicked...")
        startCopyBtn.enabled = true
        cancelCopyBtn.enabled = false
      }
    }

}
