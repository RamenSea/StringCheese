package test.test.t.test123;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.Button;

public class MainActivity extends AppCompatActivity implements View.OnClickListener {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Button button = (Button) findViewById(R.id.button);
        button.setOnClickListener(this);
    }

    int i = 0;
    @Override
    public void onClick(View v) {
        ++i;
        Button button = (Button) findViewById(R.id.button);
        button.setText(String.valueOf(i));
    }
}
