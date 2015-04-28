var showApp = angular.module("showApp", [])

showApp.controller("ServiceShow", function($scope, $http){
    $scope.plaintext = "";
    $scope.getPlainText = function(){
        console.log("I am clicking ");
        $http.get("/plain/?q=" + $scope.city).success(function(data){
            console.log(data);
            $scope.plaintext = data;
        });
    };
    $scope.getCipherText = function(){
        $http.get("/encrypted/?q=" + $scope.city).success(function(data){
            $scope.encryptedText = data;
        })
    }
    $scope.decrypted = function(){
        $http.get("/decrypted/?cipher=" + $scope.encryptedText + "&key=" + $scope.key).success(function(data){
            $scope.decryptedText = data;
        })
    }
})
