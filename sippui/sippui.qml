import QtQuick 2.2
import QtQuick.Controls 1.4
import QtQuick.Dialogs 1.0
import QtQuick.Window 2.1

ApplicationWindow {
	id: app

    title: "SIPP"

    // Hack to make the top window invisible, hopefully temporary until I can
    // figure out how to properly do SDI with QtQuick.
    width: 1
    height: 1

    menuBar: MenuBar {
        Menu {
            title: "File"
            MenuItem { 
            	text: "New Tree" 
            	shortcut: StandardKey.New
            	objectName: "newTree"
            	enabled: true
            }
            //MenuItem {
            //	text: "Open Tree"
            //	shortcut: StandardKey.Open
            //	objectName: "openTree"
            //	enabled: false
            //}
            //MenuItem {
            //	text: "Save Tree"
            //	objectName: "saveTree"
            //	enabled: false
            //}
            //MenuItem {
            //	text: "Save Image"
            //	objectName: "saveImage"
            //	enabled: false
            //}
            MenuItem {
            	text: "Close Tree (and exit)"
                shortcut: StandardKey.Close
            	objectName: "closeTree"
            	onTriggered: {
            		// Really needs an "Are you sure?" dialog
            		app.close()
            	}
            	enabled: true
            }
			enabled: true
        }
    }		
}