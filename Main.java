import java.util.Base64;

public class Main {
    public static void main(String args[]){
        String secret = "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K";
        byte[] decodeByte = Base64.getDecoder().decode(secret);
        String decodedString = new String(decodeByte);
        decodedString = decodedString.replace(":", " ");
        System.out.println("Answer: " + new StringBuilder(decodedString).reverse().toString()); //Answer: Join us at LINE MAN Wongnai;
    }
}
