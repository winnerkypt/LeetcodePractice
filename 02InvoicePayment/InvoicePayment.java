package InvoicePayment;

public class InvoicePayment{
    public static void main(String[] args) {
        // int[] invoice = {100,200,300};
        // int[] payment = {100,200,300};

        // int[] invoice = {600};
        // int[] payment = {100,200};
                        //100
        int[] invoice = {200,100,300};
        int[] payment = {100,300,200};
                        //0  200

        
        for (int i = 0; i < invoice.length; i++) {
            for (int j = 0; j < payment.length; j++){
                if(invoice[i] != 0 && payment[j] != 0) {
                    
                    if(invoice[i]<payment[j]){
                        payment[j] = payment[j]-invoice[i];
                        System.out.println("invoice "+i+" - payment "+j+ " - "+invoice[i]);
                        invoice[i] =0;
                    }else{
                        invoice[i] = invoice[i]-payment[j];
                        System.out.println("invoice "+i+" - payment "+j+ " - "+payment[j]);
                        payment[j] = 0;
                    } 
                }
            }    
            if(invoice[i]!=0){
                System.out.println("invoice "+i+" - Outstanding balance "+invoice[i]);
            }        
        }
    }
}