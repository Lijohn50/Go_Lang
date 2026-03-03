// Online C++ compiler to run C++ program online
#include <iostream>

using namespace std;

int main() {
    
    int coin[] = {1, 3, 5};
    int amount = 8;
    int sol[(sizeof(coin) / sizeof(coin[0]))+1][amount+1];
    
    for(int i = 0; i < sizeof(coin) / sizeof(coin[0]); i++){
        
        for(int j = 0; j <= amount; j++){
            
            if(i == 0 && j == 0){
                
                sol[i][j] = 1;
            }else{
                
                if(coin[i-1] < j){
                    
                    sol[i][j] = sol[i-1][j];
                }else{
                    
                    sol[i][j] = sol[i - 1][j] + sol[i][j - coin[i-1]];
                }
            }
        }
    }
    
    cout << sol[3][8];
    return 0;

}
